package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang-store/internal/config"
	"time"
)

type Claims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWTMethod struct {
	secret []byte
}

func NewJWTMethod(cfg *config.Config) *JWTMethod {
	return &JWTMethod{[]byte(cfg.JWTSecret)}
}

func (m *JWTMethod) GenerateJWT(userID int64, email, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(m.secret)
}

func (m *JWTMethod) ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", t.Header["alg"])
		}
		return m.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, fmt.Errorf("срок действия токена истек: %w", err)
		}
		return nil, fmt.Errorf("ошибка парсинга токена: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("недействительный токен")
	}

	return claims, nil
}
