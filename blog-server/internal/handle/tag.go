package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Tag struct{}

type AddOrEditTag struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

func (*Tag) CreateOrUpdateTag(c *gin.Context) {
	var tagParam AddOrEditTag

	if err := c.ShouldBindJSON(&tagParam); err != nil {
		ReturnError(c, g.ErrParamKey, err)
		return
	}

	db := model.GetDB(c)

	tag, err := model.CreateOrUpdateTag(db, tagParam.ID, tagParam.Name)

	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, tag)
}

func (*Tag) DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		ReturnError(c, g.ErrParamKey, err)
		return
	}

	db := model.GetDB(c)

	err = model.DeleteTag(db, id)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
	}
	ReturnSuccess(c, nil)
}

type TagListParam struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	Keyword  string `json:"keyword"`
}

type TagListResult struct {
	model.TagVO
}

func (*Tag) GetTagList(c *gin.Context) {
	var tagListParam TagListParam
	err := c.ShouldBindQuery(&tagListParam)

	if err != nil {
		ReturnError(c, g.ErrParamKey, err)
		return
	}

	db := model.GetDB(c)

	results, total, err := model.GetTagList(db, tagListParam.PageNum, tagListParam.PageSize, tagListParam.Keyword)

	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
	}

	lists := make([]TagListResult, 0)

	for _, list := range results {
		lists = append(lists, TagListResult{list})
	}

	ReturnSuccess(c, TagResult[TagListResult]{
		Results: lists,
		Total:   total,
	})
}
