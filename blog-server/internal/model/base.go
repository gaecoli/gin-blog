package model

import "gorm.io/gorm"

var DB *gorm.DB

// InitDB init db, at now, we just support mysql
func InitDB(db *gorm.DB) {

}
