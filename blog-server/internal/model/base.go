package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var DB *gorm.DB

// InitDB 连接数据库，目前暂时只支持 MySQL
func InitDB() (err error) {
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/gin_blog?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	} else {
		DB = db
		// TODO: add auto migrate
		return nil
	}

}

// CloseDB 关闭数据库
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}
