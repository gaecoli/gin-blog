package handle

import (
	"errors"
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sagikazarmark/slog-shim"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTP code + 业务码 + 消息 + 数据
func ReturnHttpResponse(c *gin.Context, httpCode int, code int, msg string, data any) {
	c.JSON(httpCode, Response{
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
		Response{
			Code:    r.Code(),
			Message: r.Msg(),
			Data:    data,
		},
	)
}

func ReturnSuccess(c *gin.Context, data any) {
	ReturnResponse(c, g.OkResult, data)
}

type PageResult struct {
	PageNum  int         `json:"page_num"`     // 每页条数
	PageSize int         `json:"page_size"`    // 上次页数
	Total    int64       `json:"total"`        // 总条数
	Results  interface{} `json:"page_results"` // 分页数据
}

type CategoryResult struct {
	Results interface{} `json:"results"`
	Total   int64       `json:"total"`
}

type TagResult struct {
	Results interface{} `json:"results"`
	Total   int64       `json:"total"`
}

func GetCurrentUser(c *gin.Context) (*model.User, error) {
	session := sessions.Default(c)

	id := session.Get(g.CTX_USER)
	if id == nil {
		return nil, errors.New("session 中没有找到 user id")
	}

	db := model.GetDB(c)

	// session Get 返回类型为 interface{}
	user, err := model.GetUserInfoById(db, id.(int))

	return user, err
}
