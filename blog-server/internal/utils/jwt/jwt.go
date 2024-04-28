package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt struct {
	JwtKey []byte
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(secret, email string, expire int) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(expire))
	claims := Claims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    "gin-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (*Claims, error) {
	var claims = new(Claims) // new 为 type new(Type) *Type
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	// TODO: 根据 jwt 中的错误消息，细化 token 错误

	return nil, err
}
