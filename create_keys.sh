set -e

# Keys oluştur (PKCS#1 / RSA PRIVATE KEY)
openssl genrsa -traditional -out ./api-gateway/pkg/token/keys/private_key.pem 2048
openssl rsa -in ./api-gateway/pkg/token/keys/private_key.pem -pubout -out ./api-gateway/pkg/token/keys/public_key.pem
