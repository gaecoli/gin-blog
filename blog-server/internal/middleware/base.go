package middleware

import (
	g "gin-blog/internal/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
