package model

// Category 分类
type Category struct {
	Model

	Name    string    `gorm:"unique;type:varchar(20);not null" json:"name"`
	Article []Article `gorm:"foreignKey:CategoryId"` // 一个分类下可以有多篇文章
}
