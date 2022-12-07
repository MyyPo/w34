package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	pathToPrivateSignature string
	pathToPublicSignature  string
	accessTokenDuraion     time.Duration
	refreshTokenDuration   time.Duration
}

func NewJWTManager(
	pathToPrivateSignature, pathToPublicSignature string,
	accessTokenDuration, refreshTokenDuration time.Duration,
) *JWTManager {
	return &JWTManager{
		pathToPrivateSignature: pathToPrivateSignature,
		pathToPublicSignature:  pathToPublicSignature,
		accessTokenDuraion:     accessTokenDuration,
		refreshTokenDuration:   refreshTokenDuration,
	}
}

func (m JWTManager) GenerateJWT(userUUID string) (string, error) {
	rsaPrivateSignature, err := LoadRSAPrivateKeyFromDisk(m.pathToPrivateSignature)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(m.accessTokenDuraion).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["sub"] = userUUID

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to token: %w", err)
	}

	return token, nil
}

func (m JWTManager) ValidateJwtExtractClaims(jwtTokenString string) (jwt.MapClaims, error) {
	rsaPublicSignature, err := LoadRSAPublicKeyFromDisk("../../../configs/rsa.pub")
	if err != nil {
		return nil, fmt.Errorf("failed to load the signature: %q", err)
	}

	jwtToken, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		// check if the signing algorithm is correct
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// check the signature of the token
		return rsaPublicSignature, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify: %q", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims error: %q", err)
	}

	return claims, nil
}

func LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
