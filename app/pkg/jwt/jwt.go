package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type UserClaims struct {
	Username string
	jwt.RegisteredClaims
}

type JWTManager struct {
	accessToken  string
	refreshToken string
	accessTTL    time.Duration
	refreshTTL   time.Duration
}

func NewJWTManager(accessKey, refreshKey string, accessTTL, refreshTTl time.Duration) *JWTManager {
	return &JWTManager{
		accessToken:  accessKey,
		refreshToken: refreshKey,
		accessTTL:    accessTTL,
		refreshTTL:   refreshTTl,
	}
}

func (jm *JWTManager) GenerateAccessToken(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jm.accessTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jm.accessToken))
	if err != nil {
		return "", errors.New("Error getting Access Token")
	}
	return signedToken, nil
}

func (jm *JWTManager) GenerateRefreshToken(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jm.refreshTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jm.refreshToken))
	if err != nil {
		return "", errors.New("Error getting Refresh Token")
	}
	return signedToken, nil
}

func (jm *JWTManager) VerifyToken(inputToken, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(inputToken, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("Can't parse token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil
}

func (jm *JWTManager) VerifyAccessToken(token string) (*UserClaims, error) {
	return jm.VerifyToken(token, jm.accessToken)
}

func (jm *JWTManager) VerifyRefreshToken(token string) (*UserClaims, error) {
	return jm.VerifyToken(token, jm.refreshToken)
}
