package handle

import (
	"gin-blog/internal/model"
	"github.com/gin-gonic/gin"
)

type UserAuth struct{}

func Login(c *gin.Context) {
	formUser := model.User{}

	_ = c.ShouldBindJSON(&formUser)

	db := model.GetDB(c)

	//user, err := model.CheckUserLogin(db, formUser.Email, formUser.Password)

}
