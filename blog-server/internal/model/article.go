package model

type Article struct {
	Model

	Title     string `gorm:"type:varchar(100);not null" json:"title"`
	Desc      string `json:"desc"`    // 文章描述, 用于文章列表展示
	Content   string `json:"content"` // 文章内容
	Image     string `json:"image"`
	Type      int    `gorm:"type:tinyint;comment:文章类型(1-原创 2-转载 3-翻译)" json:"type"`
	Status    int    `gorm:"type:tinyint;comment:文章状态(1-公开 2-私密)" json:"status"`
	IsTop     bool   `json:"is_top"`
	IsDeleted bool   `json:"is_deleted"`
	SourceUrl string `json:"source_url"`

	CategoryId int `json:"category_id"`

	Tags     []*Tag    `gorm:"many2many:article_tag;joinForeignKey:article_id" json:"tags"`
	Category *Category `gorm:"foreignKey:CategoryId" json:"category"`
}
