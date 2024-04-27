package middleware

import (
	"errors"
	g "gin-blog/internal/global"
	"gin-blog/internal/handle"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Jwt struct {
	JwtKey []byte
}

var secretKey = []byte(g.JwtKey)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var (
	TokenExpired   = errors.New("token 已过期，请重新登录！")
	TokenInValid   = errors.New("token 无效，请重新登录！")
	TokenError     = errors.New("token 错误，请重新登录！")
	HeaderNotFound = errors.New("header 没找到，请重新登录！")
)

func GenerateToken(username string) (string, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * time.Duration(g.Conf.Jwt.ExpireDays))
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    "gin-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*Claims, error) {
	var claims = new(Claims) // new 为 type new(Type) *Type
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
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

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			handle.ReturnError(ctx, g.ErrRequestLogin, HeaderNotFound)
			return
		}

		// token 的正确格式: `Bearer xxxxxxxx`
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			handle.ReturnError(ctx, g.ErrRequestLogin, TokenInValid)
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			handle.ReturnError(ctx, g.ErrRequestLogin, TokenInValid)
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
