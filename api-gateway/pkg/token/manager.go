package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	expiresIn  time.Duration
}

func NewManager(config config.TokenConfig) (*Manager, error) {

	privateKey, err := loadPrivateKey(config.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}

	publicKey, err := loadPublicKey(config.PublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create token manager: %w", err)
	}
	return &Manager{
		privateKey: privateKey,
		publicKey:  publicKey,
		expiresIn:  time.Duration(config.TokenExpiresInMinutes) * time.Minute,
	}, nil
}

func (m *Manager) GenerateAccessToken(userId string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "taxihub",
		Subject:   userId,
		ExpiresAt: jwt.NewNumericDate(now.Add(m.expiresIn)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	signedToken, err := token.SignedString(m.privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing access token: %w", err)
	}

	return signedToken, nil
}

func (m *Manager) VerifyAccessToken(accessToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return m.publicKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}),
		jwt.WithExpirationRequired(),
		jwt.WithIssuedAt(),
	)

	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return nil, fmt.Errorf("token expired: %w", err)
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return nil, fmt.Errorf("token not valid yet: %w", err)
		default:
			return nil, fmt.Errorf("invalid token: %w", err)
		}
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read private key file: %w", err)
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing RSA private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}

	// Decode PEM block
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	// Parse the public key
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Assert that the key is an RSA public key
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not an RSA public key")
	}

	return rsaPublicKey, nil
}
