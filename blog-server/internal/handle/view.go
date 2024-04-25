package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"github.com/gin-gonic/gin"
)

type BlogInfo struct{}

type BlogHomeVO struct {
	ArticleCount int `json:"article_count"`
}

type AboutReq struct {
	Content string `json:"content"`
}

func (*BlogInfo) GetConfigMap(c *gin.Context) {
	db := model.GetDB(c)

	data, err := model.GetConfigMap(db)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, data)
}

func (*BlogInfo) UpdateConfig(c *gin.Context) {
	//var m map[string]string

}

func (*BlogInfo) GetHomeInfo(c *gin.Context) {
	db := model.GetDB(c)

	articleCount, err := model.Count(db, &model.Article{}, "status = ? AND is_deleted = ?", 1, 0)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	// TODO: 返回 home info, 暂时返回文章的总数
	ReturnSuccess(c, BlogHomeVO{
		ArticleCount: articleCount,
	})
}
