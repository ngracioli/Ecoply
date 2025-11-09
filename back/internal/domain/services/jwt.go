package services

import (
	"ecoply/internal/config"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
)

type Claims struct {
	UserUuid  string `json:"user_uuid"`
	UserEmail string `json:"user_email"`
	UserType  string `json:"user_type"`
	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateToken(userUuid string, email string, userType string) (string, error)
	RefreshToken(oldToken string) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type jwtService struct {
	signingKey []byte
	issuer     string
}

func NewJwtService(cfg *config.Config) JwtService {
	return &jwtService{
		signingKey: []byte(cfg.JWTSigningKey),
		issuer:     cfg.AppName,
	}
}

func (j *jwtService) GenerateToken(userUuid string, email string, userType string) (string, error) {
	now := time.Now()
	expirationTime := now.Add(24 * time.Hour)

	claims := &Claims{
		UserUuid:  userUuid,
		UserEmail: email,
		UserType:  userType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    j.issuer,
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(j.signingKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (j *jwtService) RefreshToken(oldToken string) (string, error) {
	claims, err := j.ValidateToken(oldToken)
	if err != nil {
		return "", err
	}

	if time.Until(claims.ExpiresAt.Time) > 5*time.Minute {
		return oldToken, nil
	}

	newToken, err := j.GenerateToken(claims.UserUuid, claims.UserEmail, claims.UserType)
	if err != nil {
		return "", err
	}

	return newToken, nil
}
