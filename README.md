# TaxiHub

1) Public ve private keylerinizi oluşturun:

```bash
chmod +x ./create_keys.sh
./create_keys.sh
```

Eğer keylerinizi manuel olarak oluşturmak isterseniz lütfen buradaki dosyalara kaydedin.
```bash
./api-gateway/pkg/token/keys/private_key.pem
./api-gateway/pkg/token/keys/public_key.pem
```

2) Tüm servisleri Docker Compose ile ayağa kaldırın:

```bash
docker compose up --build
```

3) JWT Token Oluşturma

Bu endpointi kullanarak protected route'ları test etmek için bir JWT token alabilirsiniz. (Sadece test içindir)
```bash
GET http://localhost:8080/api/token
```

4) API Kullanımı

Swagger dökümantasyonuna bu URL'den ulaşabilirsiniz. (NOT: JWT Auth gerektiren endpointler için Swagger'da `Authorize` alanına `Bearer <token>` formatında girin.)
```bash
(http://localhost:8080/swagger/index.html)
```

## Notlar

- Config dosyaları hızlıca test etmek için repoya push edilmiştir. Production için .gitignore'a eklenip server tarafında config dosyaları oluşturulmalı veya secret management toolları kullanılmalıdır.
