package services

import (
	"ecoply/internal/config"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

type JWTService struct {
	signingKey []byte
	issuer     string
}

func NewJWTService() *JWTService {
	cfg := config.GetConfig()
	return &JWTService{
		signingKey: []byte(cfg.JWTSigningKey),
		issuer:     cfg.AppName,
	}
}

func (j *JWTService) GenerateToken(userID uint, email string, userType string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Email:    email,
		UserType: userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    j.issuer,
			Subject:   email,
			ID:        string(rune(userID)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// TODO Implementar rota de refresh
func (j *JWTService) RefreshToken(tokenString string) (string, error) {
	claims, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if time.Until(claims.ExpiresAt.Time) > time.Hour {
		return tokenString, nil
	}

	return j.GenerateToken(claims.UserID, claims.Email, claims.UserType)
}

func ExtractTokenFromBearer(authHeader string) string {
	var token []string = strings.Split(authHeader, "Bearer ")
	if len(token) == 2 {
		return token[1]
	}
	return ""
}
