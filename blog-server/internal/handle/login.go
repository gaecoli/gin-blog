package handle

import (
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"gin-blog/internal/utils/jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

type LoginResponse struct {
	user  model.User
	token string
}

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (*LoginApi) Login(c *gin.Context) {
	var formUser LoginInfo

	_ = c.ShouldBindJSON(&formUser)

	db := model.GetDB(c)

	user, err := model.CheckUserLogin(db, formUser.Email, formUser.Password)

	if err != nil {
		ReturnError(c, g.ErrRequestLogin, err)
		return
	}

	token, err := jwt.GenerateToken(g.Conf.Jwt.JwtKey, user.Email, g.Conf.Jwt.ExpireDays)

	if err != nil {
		ReturnError(c, g.ErrRequestLogin, err)
		return
	}

	session := sessions.Default(c)
	session.Set(g.CTX_USER, user.ID)
	sessionErr := session.Save()
	if sessionErr != nil {
		ReturnError(c, g.ErrRequestLogin, err)
	}

	ReturnSuccess(c, LoginResponse{
		user:  user,
		token: token,
	})
}

func (*LoginApi) Logout(c *gin.Context) {

	session := sessions.Default(c)
	session.Delete(g.CTX_USER)
	err := session.Save()
	if err != nil {
		return
	}

	ReturnSuccess(c, nil)
}

func (*LoginApi) Register(c *gin.Context) {
	// 创建一个 User 结构体来接收用户提交的注册信息
	newUser := model.User{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		ReturnError(c, g.ErrRequestRegister, err)
		return
	}

	// 对用户的密码进行加密
	hashedPassword, err := model.PasswordHashString(newUser.Password)
	if err != nil {
		ReturnError(c, g.ErrRequestRegister, err)
		return
	}
	newUser.Password = hashedPassword

	// 将用户的信息保存到数据库中
	db := model.GetDB(c)
	userErr := model.CreateUserInfo(db, &newUser)
	if userErr != nil {
		ReturnError(c, g.ErrRequestRegister, userErr)
	}

	// 返回成功的响应
	ReturnSuccess(c, "注册用户成功")
}

func (*LoginApi) SendEmailCode(c *gin.Context) {
	ReturnSuccess(c, "发送邮箱验证码")
}
