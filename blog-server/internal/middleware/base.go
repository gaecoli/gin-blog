package middleware

import (
	g "gin-blog/internal/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WithGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(g.CTX_DB, db)
		// TODO: 分析一下 Next() 源码
		ctx.Next() // 将控制权交给 gin，否则请求将在中间件函数中停止，不继续执行后续中间件或者路由处理函数
	}
}
