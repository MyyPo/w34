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
	accessPrivateSignature  *rsa.PrivateKey
	AccessPublicSignature   *rsa.PublicKey
	refreshPrivateSignature *rsa.PrivateKey
	RefreshPublicSignature  *rsa.PublicKey
	accessTokenDuraion      time.Duration
	refreshTokenDuration    time.Duration
}

func NewJWTManager(
	pathToAccessPrivateSignature, pathToAccessPublicSignature string,
	pathToRefreshPrivateSignature, pathToRefreshPublicSignature string,
	accessTokenDuration, refreshTokenDuration time.Duration,
) (*JWTManager, error) {
	accessPrivateSignature, err := loadRSAPrivateKeyFromDisk(pathToAccessPrivateSignature)
	if err != nil {
		return nil, err
	}
	accessPublicSignature, err := loadRSAPublicKeyFromDisk(pathToAccessPublicSignature)
	if err != nil {
		return nil, err
	}
	refreshPrivateSignature, err := loadRSAPrivateKeyFromDisk(pathToRefreshPrivateSignature)
	if err != nil {
		return nil, err
	}
	refreshPublicSignature, err := loadRSAPublicKeyFromDisk(pathToRefreshPublicSignature)
	if err != nil {
		return nil, err
	}

	return &JWTManager{
		accessPrivateSignature:  accessPrivateSignature,
		AccessPublicSignature:   accessPublicSignature,
		refreshPrivateSignature: refreshPrivateSignature,
		RefreshPublicSignature:  refreshPublicSignature,
		accessTokenDuraion:      accessTokenDuration,
		refreshTokenDuration:    refreshTokenDuration,
	}, nil
}

type Claims struct {
	TknType string `json:"tkn_type"`
	jwt.RegisteredClaims
}

func (m JWTManager) GenerateAccessToken(userID int32) (string, error) {

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

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(m.accessPrivateSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to access token: %w", err)
	}

	return accessToken, nil
}

func (m JWTManager) GenerateRefreshToken(userID int32) (string, error) {

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

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(m.refreshPrivateSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to refresh token: %w", err)
	}

	return refreshToken, nil
}

func (m JWTManager) ValidateJwtExtractClaims(jwtTokenString string, publicSignature *rsa.PublicKey) (*Claims, error) {

	var claims *Claims
	jwtToken, err := jwt.ParseWithClaims(jwtTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// check if the signing algorithm is correct
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// check the signature of the token
		return publicSignature, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to verify: %q", err)
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("claims error")
	}

	return claims, nil
}

func loadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
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

func loadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
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
