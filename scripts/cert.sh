#!/bin/bash

openssl req -x509 -newkey rsa:2048 -keyout okta-app.key -out okta-app.cert -days 365 -nodes -subj "/CN=example.com"

echo "Certificate generated successfully!"