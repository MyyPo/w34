package jwt

import (
	"crypto/rsa"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	pathToAccessPrivateSignature  string
	pathToAccessPublicSignature   string
	pathToRefreshPrivateSignature string
	PathToRefreshPublicSignature  string
	accessTokenDuraion            time.Duration
	refreshTokenDuration          time.Duration
}

func NewJWTManager(
	pathToAccessPrivateSignature, pathToAccessPublicSignature string,
	pathToRefreshPrivateSignature, pathToRefreshPublicSignature string,
	accessTokenDuration, refreshTokenDuration time.Duration,
) *JWTManager {
	return &JWTManager{
		pathToAccessPrivateSignature:  pathToAccessPrivateSignature,
		pathToAccessPublicSignature:   pathToAccessPublicSignature,
		pathToRefreshPrivateSignature: pathToRefreshPrivateSignature,
		PathToRefreshPublicSignature:  pathToRefreshPublicSignature,
		accessTokenDuraion:            accessTokenDuration,
		refreshTokenDuration:          refreshTokenDuration,
	}
}

type Claims struct {
	TknType string `json:"tkn_type"`
	jwt.RegisteredClaims
}

func (m JWTManager) GenerateAccessToken(userID int32) (string, error) {
	rsaPrivateAccessSignature, err := m.LoadRSAPrivateKeyFromDisk(m.pathToAccessPrivateSignature)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	expires := now.Add(m.accessTokenDuraion)
	numericDateNow := jwt.NewNumericDate(now)
	numericDateExpires := jwt.NewNumericDate(expires)

	strUserID := strconv.FormatInt(int64(userID), 10)

	claims := Claims{
		"access",
		jwt.RegisteredClaims{
			ExpiresAt: numericDateExpires,
			IssuedAt:  numericDateNow,
			NotBefore: numericDateNow,
			Subject:   strUserID,
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateAccessSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to access token: %w", err)
	}

	return accessToken, nil
}

func (m JWTManager) GenerateRefreshToken(userID int32) (string, error) {
	rsaPrivateRefreshSignature, err := m.LoadRSAPrivateKeyFromDisk(m.pathToRefreshPrivateSignature)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()
	expires := now.Add(m.refreshTokenDuration)
	numericDateNow := jwt.NewNumericDate(now)
	numericDateExpires := jwt.NewNumericDate(expires)

	strUserID := strconv.FormatInt(int64(userID), 10)

	claims := Claims{
		"refresh",
		jwt.RegisteredClaims{
			ExpiresAt: numericDateExpires,
			IssuedAt:  numericDateNow,
			NotBefore: numericDateNow,
			Subject:   strUserID,
		},
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateRefreshSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to refresh token: %w", err)
	}

	return refreshToken, nil
}

func (m JWTManager) ValidateJwtExtractClaims(jwtTokenString, publicSignaturePath string) (*Claims, error) {
	rsaPublicSignature, err := m.LoadRSAPublicKeyFromDisk(publicSignaturePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load the signature: %q", err)
	}

	var claims *Claims
	jwtToken, err := jwt.ParseWithClaims(jwtTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("claims error: %q", err)
	}

	return claims, nil
}

func (m JWTManager) LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
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

func (m JWTManager) LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
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