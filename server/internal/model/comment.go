package model

// 暂不开放
type Comment struct {
	Model

	Message string `json:"message"`

	UserId int   `json:"user_id"`
	User   *User `gorm:"foreignKey:UserId" json:"-"`
}
