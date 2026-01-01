# TaxiHub

TaxiHub, API Gateway + Driver Service + MongoDB'den olusan basit bir mikroservis ornegidir.

## Hizli baslangic (tek komutla calistirma)

1) Token key'lerini uretin:

```bash
chmod +x ./create_keys.sh
./create_keys.sh
```

2) Tum servisi Docker Compose ile kaldirin:

```bash
docker compose up --build
```

Hepsi bu kadar. Servisler ayaga kalktiktan sonra API Gateway su adreslerden kullanilabilir:

- Swagger UI: `http://localhost:8080/swagger/index.html`
- Health: `http://localhost:8080/health`
- Token: `http://localhost:8080/token`
- API BasePath (auth gerekli endpoint'ler): `http://localhost:8080/api`

## Gereksinimler

- Docker + Docker Compose
- OpenSSL (token key'lerini uretmek icin)

## Notlar

- `api-gateway/internal/config/config.yaml` ve `driver-service/internal/config/config.yaml` Docker icinde kullanilan ayarlari icerir.
- Driver endpoint'leri auth gerektirir. Swagger'da `Authorize` alanina `Bearer <token>` formatinda girin.
