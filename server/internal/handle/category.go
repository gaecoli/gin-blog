package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Category struct{}

func (*Category) GetCategoryList(c *gin.Context) {
	db := model.GetDB(c)

	results, total, err := model.GetCategories(db)

	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, CategoryResult{
		Results: results,
		Total:   total,
	})
}

func (*Category) CreateCategory(c *gin.Context) {
	db := model.GetDB(c)
	data := model.Category{}
	_ = c.ShouldBindJSON(&data)
	err := model.CreateCategory(db, &data)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, data)
}

func (*Category) UpdateCategory(c *gin.Context) {
	db := model.GetDB(c)
	data := model.Category{}
	_ = c.ShouldBindJSON(&data)
	err := model.UpdateCategory(db, &data)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, data)
}

func (*Category) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id == 0 {
		ReturnError(c, g.ErrParamKey, err)
		return
	}
	db := model.GetDB(c)
	err = model.DeleteCategory(db, id)
	if err != nil {
		ReturnError(c, g.ErrDbOp, err)
		return
	}
	ReturnSuccess(c, nil)
}
