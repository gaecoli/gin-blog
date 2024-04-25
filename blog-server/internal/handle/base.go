package handle

import (
	g "gin-blog/internal/global"
	"github.com/gin-gonic/gin"
	"github.com/sagikazarmark/slog-shim"
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

// TODO: 学习 error 处理
func ReturnError(c *gin.Context, r g.Result, data any) {
	slog.Info("[Func-ReturnError]" + r.Msg())

	val := r.Msg()

	if data != nil {
		switch v := data.(type) {
		case error:
			val = v.Error()
		case string:
			val = v
		}
		slog.Error(val) // 错误日志
	}

	c.AbortWithStatusJSON(
		http.StatusOK,
		Response[any]{
			Code:    r.Code(),
			Message: r.Msg(),
			Data:    data,
		},
	)
}

func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, g.OkResult, data)
}
