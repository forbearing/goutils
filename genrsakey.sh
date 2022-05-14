#!/usr/bin/env bash

openssl genrsa -out private.pem 1024
openssl rsa -in private.pem -pubout -out public.pem
openssl pkcs8 -topk8 -inform pem -in private.pem -outform PEM -nocrypt > public.pem.pkcs8
ln -sf private.pem private.pem.pkcs1
