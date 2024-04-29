package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Article struct{}

func (*Article) GetArticle(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	db := model.GetDB(c)

	article, err := model.GetArt(db, id)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	if article == nil {
		ReturnError(c, g.ErrNotFound, nil)
		return
	}

	ReturnSuccess(c, article)
}

func (*Article) CreateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	currentUser, err := GetCurrentUser(c)

	if err != nil || currentUser == nil {
		ReturnError(c, g.ErrNotFound, "登录错误，不可以编辑文章")
		return
	}

	db := model.GetDB(c)

	data.UserId = currentUser.ID

	article := model.CreateArt(db, &data)

	ReturnSuccess(c, article)
}

func (*Article) UpdateArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	db := model.GetDB(c)

	article := model.UpdateArt(db, &data)

	ReturnSuccess(c, article)
}

func (*Article) DeleteArticle(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	db := model.GetDB(c)

	err = model.DeleteArt(db, id)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, nil)
}

type SoftArchiveArticleParam struct {
	ID int `json:"id" binding:"required"`
}

func (*Article) SoftDeleteArticle(c *gin.Context) {
	var archiveParam SoftArchiveArticleParam
	err := c.ShouldBindJSON(&archiveParam)
	if err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	db := model.GetDB(c)

	_, err = model.SoftDeleteArt(db, archiveParam.ID)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, nil)
}

type ArticlePageParam struct {
	PageNum  int `form:"page_num"`
	PageSize int `form:"page_size"`
}

func (*Article) GetArticleList(c *gin.Context) {
	var query ArticlePageParam
	err := c.ShouldBindQuery(&query)
	if err != nil {
		ReturnError(c, g.ErrRequest, query)
		return
	}

	db := model.GetDB(c)

	articles, total, err := model.GetArticleList(db, query.PageNum, query.PageSize)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, PageResult{
		PageNum:  query.PageNum,
		PageSize: query.PageSize,
		Total:    total,
		Results:  articles,
	})

}
