package model

import (
	"errors"
	util "gin-blog/internal/utils"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	Model

	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Name      string    `gorm:"unique;type:varchar(320);not null" json:"name"`
	Avatar    string    `gorm:"type:varchar(320)" json:"avatar"`
	Password  string    `gorm:"type:varchar(320)" json:"-"`
	DisableAt time.Time `gorm:"default:null" json:"disable_at"`
	IsAdmin   bool      `json:"is_admin"`
}

func GetUserInfoById(db *gorm.DB, id int) (user *User, err error) {
	err = db.Model(&User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserInfoByName(db *gorm.DB, name string) (user *User, err error) {
	err = db.Model(&User{}).Where("name = ?", name).First(&user).Error
	return user, err
}

func UpdateUserInfo(db *gorm.DB, id int, email, name, avatar string) error {
	userInfo := User{
		Model:  Model{ID: id},
		Email:  email,
		Avatar: avatar,
		Name:   name,
	}

	err := db.Model(&User{}).Updates(&userInfo).Error

	return err
}

func CreateUserInfo(db *gorm.DB, user *User) error {
	if user.ID == 0 {
		hashedPassword, err := PasswordHashString(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
		
		parts := strings.Split(user.Email, "@")
		if len(parts) < 1 {
			return errors.New("邮箱错误")
		}
		name := parts[0]

		user.Name = name

		err = db.Model(&User{}).Create(&user).Error
		return err
	}

	return errors.New("创建用户错误")
}

func UpdateUserPassword(db *gorm.DB, id int, password string) error {
	hashedPassword, err := PasswordHashString(password)
	if err != nil {
		return err
	}

	err = db.Model(&User{}).Where("id = ?", id).Update("password", hashedPassword).Error
	return err
}

func UpdateUserDisableAt(db *gorm.DB, id int, disable bool) (err error) {
	if disable {
		disableAt := time.Now()

		err = db.Model(&User{}).Where("id = ?", id).Update("disabled_at", disableAt).Error
	} else {
		err = db.Model(&User{}).Where("id = ?", id).Update("disabled_at", nil).Error
	}

	return err
}

func CheckUserLogin(db *gorm.DB, email string, password string) (User, error) {
	var user User

	if !util.IsValidEmail(email) {
		return user, errors.New("无效邮箱")
	}

	passwdErr := util.IsValidPassword(password)

	if passwdErr != nil {
		return user, passwdErr
	}

	err := db.Where("email = ?", email).First(&user).Error

	if user.ID == 0 {
		return user, errors.New("未找到用户！")
	}

	verify := PasswordVerify(user.Password, password)

	if !verify {
		return user, err
	}

	// TODO: 后续添加校验用户角色

	return user, nil
}

// 后续也不用校验用户角色
func CheckUserFrontLogin(db *gorm.DB, email string, password string) (User, error) {
	var user User

	if !util.IsValidEmail(email) {
		return user, errors.New("用户邮箱无效")
	}

	passwdErr := util.IsValidPassword(password)

	if passwdErr != nil {
		return user, passwdErr
	}

	verify := PasswordVerify(user.Password, password)

	if !verify {
		return user, errors.New("密码不正确")
	}

	err := db.Where("email = ?", email).First(&user).Error

	return user, err
}
