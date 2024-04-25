package model

import "gorm.io/gorm"

type Article struct {
	Model

	Title     string  `gorm:"type:varchar(100);not null" json:"title"`
	Desc      string  `json:"desc"`    // 文章描述, 用于文章列表展示
	Content   string  `json:"content"` // 文章内容
	Image     *string `json:"image"`
	Type      int     `gorm:"type:tinyint;comment:文章类型(1-原创 2-转载 3-翻译)" json:"type"`
	Status    int     `gorm:"type:tinyint;comment:文章状态(1-公开 2-私密)" json:"status"`
	IsTop     bool    `json:"is_top"`
	IsDeleted bool    `json:"is_deleted"`
	SourceUrl string  `json:"source_url"`

	CategoryId int `json:"category_id"`

	Tags     []*Tag    `gorm:"many2many:article_tag;joinForeignKey:article_id" json:"tags"`
	Category *Category `gorm:"foreignKey:CategoryId" json:"category"`
}

// 新建文章
func CreateArt(db *gorm.DB, article *Article) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// TODO: 需要补充 tags 和 category 的处理
		result := db.Create(&article)
		if result.Error != nil {
			return result.Error
		}
		return result.Error
	})
}

func UpdateArt(db *gorm.DB, article *Article) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var result *gorm.DB
		if article.ID == 0 {
			return nil
		} else {
			result = db.Model(&article).Where("id = ?", article.ID).Updates(article)
		}
		if result.Error != nil {
			return result.Error
		}
		return result.Error
	})
}

func GetArt(db *gorm.DB, id int) (data *Article, err error) {
	if id == 0 {
		return nil, nil
	}

	result := db.Model(&Article{}).Where("id = ?", id).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, result.Error
}

func DeleteArt(db *gorm.DB, id int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		result := db.Where("id = ?", id).Delete(&Article{})
		if result.Error != nil {
			return result.Error
		}
		return result.Error
	})
}

func SoftDeleteArt(db *gorm.DB, id int) (int64, error) {
	result := db.Model(&Article{}).Where("id = ?", id).Update("is_deleted", true)
	if result.Error != nil {
		return 0, result.Error
	}
	return 1, nil
}
