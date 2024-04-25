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

	db := model.GetDB(c)

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

func (*Article) SoftDeleteArticle(c *gin.Context) {
	stringId := c.Param("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		ReturnError(c, g.ErrRequest, err)
		return
	}

	db := model.GetDB(c)

	_, err = model.SoftDeleteArt(db, id)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, nil)
}
