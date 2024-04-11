package handle

import (
	g "gin-blog/internal/global"
	"github.com/gin-gonic/gin"
	"github.com/sagikazarmark/slog-shim"
	"gorm.io/gorm"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

// HTTP code + 业务码 + 消息 + 数据
func ReturnHttpResponse(c *gin.Context, httpCode int, code int, msg string, data any) {
	c.JSON(httpCode, Response[any]{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func ReturnResponse(c *gin.Context, r g.Result, data any) {
	ReturnHttpResponse(c, http.StatusOK, r.Code(), r.Msg(), data)
}

func GetDB(c *gin.Context) *gorm.DB {
	// gin MustGet返回一个 any (interface{})
	// 通过类型断言的方式将它转为想要的 *gorm.DB 类型
	// 例子：
	// var a interface{} = 10
	// t, ok := a.(int) 转为整数类型
	// t1, ok1 := a.(float32) 转为浮点数类型
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}

func ReturnError(c *gin.Context, r g.Result, data any) {
	slog.Info("[Func-ReturnError]" + r.Msg())
}

func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, g.OkResult, data)
}
