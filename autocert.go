package log4shell

import (
	"crypto/tls"
	"os"

	"github.com/pkg/errors"
	"golang.org/x/crypto/acme/autocert"
)

func TestAutoCert(domain string) (*tls.Certificate, error) {
	const certDir = "autocert"

	err := os.MkdirAll(certDir, 0700)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	mgr := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache(certDir),
	}
	clientHello := tls.ClientHelloInfo{
		ServerName: domain,
	}
	tlsCert, err := mgr.GetCertificate(&clientHello)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign certificate")
	}
	return tlsCert, nil
}
