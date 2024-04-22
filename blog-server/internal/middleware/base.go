package middleware

import (
	g "gin-blog/internal/global"
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
