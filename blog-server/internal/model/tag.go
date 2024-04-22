package model

// gorm omitempty: 如果该字段为空值，那么 gorm 每次将自动忽略该字段

type Tag struct {
	Model

	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`

	// 反向关联该标签下的文章
	Articles []*Article `gorm:"many2many:article_tag" json:"articles,omitempty"`
}
