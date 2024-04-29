package middleware

import (
	"errors"
	g "gin-blog/internal/global"
	"gin-blog/internal/handle"
	"gin-blog/internal/utils/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

var (
	TokenInValid   = errors.New("token 无效，请重新登录！")
	TokenError     = errors.New("token 错误，请重新登录！")
	HeaderNotFound = errors.New("header 没找到，请重新登录！")
)

func SetGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(g.CTX_DB, db)
		// 必须设置，不然会阻塞后面的处理函数;
		ctx.Next()
	}
}

func WithCookieMiddle(name, secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 600,
	})
	return sessions.Sessions(name, store)
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

		claims, err := jwt.ParseToken(g.Conf.Jwt.JwtKey, parts[1])
		if err != nil {
			handle.ReturnError(ctx, g.ErrRequestLogin, TokenInValid)
		}

		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
