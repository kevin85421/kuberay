#!/usr/bin/env bash

## Create CA
openssl req -x509 \
            -sha256 -days 356 \
            -nodes \
            -newkey rsa:2048 \
            -subj "/CN=*.test.com/C=US/L=San Francisco" \
            -keyout ca.key -out ca.crt

## Create tls.key
openssl genrsa -out tls.key 2048

## Write CSR Config
cat > csr.conf <<EOF
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = US
ST = California
L = San Fransisco
O = test
OU = test
CN = *.test.com

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = localhost
DNS.2 = service-ray-head.default.svc.cluster.local
DNS.3 = service-ray-model-training-worker.default.svc.cluster.local
IP.1 = 127.0.0.1

EOF

## Create CSR using tls.key
openssl req -new -key tls.key -out ib.csr -config csr.conf

## Write cert config
cat > cert.conf <<EOF

authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
DNS.2 = service-ray-head.default.svc.cluster.local
DNS.3 = service-ray-model-training-worker.default.svc.cluster.local
IP.1 = 127.0.0.1


EOF

## Generate tls.cert
openssl x509 -req \
    -in ib.csr \
    -CA ca.crt -CAkey ca.key \
    -CAcreateserial -out tls.crt \
    -days 365 \
    -sha256 -extfile cert.conf