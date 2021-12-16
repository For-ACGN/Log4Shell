package log4j2

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func testGenerateConfig() *Config {
	return &Config{
		LogOut:         os.Stdout,
		Hostname:       "127.0.0.1",
		ClassDirectory: "testdata",
		HTTPNetwork:    "tcp",
		HTTPAddress:    "127.0.0.1:8088",
		LDAPNetwork:    "tcp",
		LDAPAddress:    "127.0.0.1:389",
	}
}

func TestLog4j2(t *testing.T) {
	server, err := New(testGenerateConfig())
	require.NoError(t, err)

	err = server.Start()
	require.NoError(t, err)

	// select {}

	err = server.Stop()
	require.NoError(t, err)
}

func TestLog4j2TLS(t *testing.T) {
	_, cert, pri := testGenerateCert(t, "127.0.0.1")
	crt, err := tls.X509KeyPair(cert, pri)
	require.NoError(t, err)

	cfg := testGenerateConfig()
	cfg.EnableTLS = true
	cfg.TLSCert = crt

	server, err := New(cfg)
	require.NoError(t, err)

	err = server.Start()
	require.NoError(t, err)

	// select {}

	err = server.Stop()
	require.NoError(t, err)
}

// TLSCertificate is used to generate CA ASN1 data, signed certificate.
func testGenerateCert(t testing.TB, ipv4 string) (caASN1 []byte, cPEMBlock, cPriPEMBlock []byte) {
	// generate CA certificate
	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(12345678),
		SubjectKeyId: []byte{1, 2, 3, 4},
		NotBefore:    time.Now().AddDate(0, 0, -1),
		NotAfter:     time.Now().AddDate(0, 0, 1),
	}
	caCert.Subject.CommonName = "test CA"
	caCert.KeyUsage = x509.KeyUsageCertSign
	caCert.BasicConstraintsValid = true
	caCert.IsCA = true
	caPri, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	caPub := &caPri.PublicKey
	caASN1, err = x509.CreateCertificate(rand.Reader, caCert, caCert, caPub, caPri)
	require.NoError(t, err)
	caCert, err = x509.ParseCertificate(caASN1)
	require.NoError(t, err)

	// sign certificate
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(4862131),
		SubjectKeyId: []byte{4, 2, 3, 4},
		NotBefore:    time.Now().AddDate(0, 0, -1),
		NotAfter:     time.Now().AddDate(0, 0, 1),
	}
	cert.Subject.CommonName = "test certificate"
	cert.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment
	cert.DNSNames = []string{"localhost"}
	cert.IPAddresses = []net.IP{net.ParseIP(ipv4), net.ParseIP("::1")}
	cPri, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	cPub := &cPri.PublicKey
	cASN1, err := x509.CreateCertificate(rand.Reader, cert, caCert, cPub, caPri)
	require.NoError(t, err)

	// encode with pem
	cPEMBlock = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cASN1,
	})
	cPriPEMBlock = pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(cPri),
	})
	return
}
