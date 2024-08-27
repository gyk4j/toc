#!/bin/bash

if [[ ! -e /root/conf/private.key ]]; then
    openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -subj "/C=US/ST=Apache/L=Fundation/O=/CN=james.apache.org" -keyout /root/conf/private.key -out /root/conf/private.csr
fi

if [[ ! -e /root/conf/keystore ]]; then
    keytool -genkeypair -alias james -keyalg RSA -storetype PKCS12 -keystore /root/conf/keystore -storepass password -dname "C=US,ST=Delaware,L=Wilmington,O=Apache Software Foundation,CN=James Apache"
fi

wait-for-it.sh --host=localhost --port=9999 --strict --timeout=0 -- ./initialdata.sh &

java -Djdk.tls.ephemeralDHKeySize=2048 \
     -classpath '/root/resources:/root/classes:/root/libs/*' \
     -javaagent:/root/libs/openjpa-3.2.0.jar \
     -Dlogback.configurationFile=/root/conf/logback.xml \
      -Dworking.directory=/root/ org.apache.james.JPAJamesServerMain
