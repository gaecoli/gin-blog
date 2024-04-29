package model

import "gorm.io/gorm"

// Category 分类
type Category struct {
	Model

	Name    string    `gorm:"unique;type:varchar(20);not null" json:"name"`
	Article []Article `gorm:"foreignKey:CategoryId"` // 一个分类下可以有多篇文章
}

type CategoryVO struct {
	Category

	ArticleCount int `json:"article_count"`
}

// GetCategories 获取分类列表
func GetCategories(db *gorm.DB) ([]CategoryVO, int64, error) {
	results := make([]CategoryVO, 0)

	var total int64

	db = db.Table("category as c").
		Select("c.id, c.name, c.created_at, c.updated_at, count(a.id) as article_count").
		Joins("left join article a on a.category_id = c.id").
		Where("a.is_deleted = 0 AND a.status = 1")

	err := db.Group("c.id").Order("c.id desc").Count(&total).Find(&results).Error
	return results, total, err
}

func CreateCategory(db *gorm.DB, category *Category) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&category).Error
		return err
	})
}

func UpdateCategory(db *gorm.DB, category *Category) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if category.ID == 0 {
			return nil
		}
		err := tx.Model(&category).Where("id = ?", category.ID).Updates(category).Error
		return err
	})
}

func DeleteCategory(db *gorm.DB, id int) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Category{}).Where("id = ?", id).Delete(&Category{}).Error
		return err
	})
}
