package handle

import (
	"gin-blog/internal/global"
	"gin-blog/internal/model"
	util "gin-blog/internal/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type User struct{}

func (*User) UpdateUserInfo(c *gin.Context) {
	var user model.User

	_ = c.ShouldBindJSON(&user)

	db := model.GetDB(c)

	err := model.UpdateUserInfo(db, user.ID, user.Email, user.Name, user.Avatar)

	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, user)
}

func (*User) GetUserInfoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ReturnError(c, global.ErrParamKey, err)
		return
	}

	db := model.GetDB(c)
	user, err := model.GetUserInfoById(db, id)

	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, user)
}

type UpdatePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (*User) UpdateUserPassword(c *gin.Context) {
	var password UpdatePasswordReq
	if err := c.ShouldBindJSON(&password); err != nil {
		ReturnError(c, global.ErrRequest, err)
		return
	}

	currentUser, err := GetCurrentUser(c)

	if err != nil {
		ReturnError(c, global.ErrFoundKeyInSession, err)
		return
	}

	if !util.CryptCheck(password.OldPassword, currentUser.Password) {
		ReturnError(c, global.ErrOldPassword, nil)
		return
	}

	db := model.GetDB(c)
	err = model.UpdateUserPassword(db, currentUser.ID, password.NewPassword)

	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	// 强制用户下线
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	ReturnSuccess(c, "密码修改正确，请重新登录！")
}

type UserIdReq struct {
	ID int `json:"id" binding:"required"`
}

func (*User) UpdateUserDisableInfo(c *gin.Context) {
	var userReq UserIdReq
	_ = c.ShouldBindJSON(&userReq)

	db := model.GetDB(c)

	err := model.UpdateUserDisableAt(db, userReq.ID)

	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, "已经更新用户状态")
}

func (*User) GetUserInfoList(c *gin.Context) {
	pageNum, pageErr := strconv.Atoi(c.Param("page_num"))
	pageSize, pageSizeErr := strconv.Atoi(c.Param("page_size"))
	keyword := c.Param("keyword")

	var err error

	if pageErr != nil || pageSizeErr != nil {
		if pageErr != nil {
			err = pageErr
		} else {
			err = pageSizeErr
		}

		ReturnError(c, global.ErrParamKey, err)
	}

	db := model.GetDB(c)

	users, total, err := model.GetUserInfoList(db, pageNum, pageSize, keyword)

	if err != nil {
		ReturnError(c, global.ErrDbOp, err)
		return
	}

	ReturnSuccess(c, PageResult{
		PageSize: pageSize,
		PageNum:  pageNum,
		Total:    total,
		Results:  users,
	})
}
