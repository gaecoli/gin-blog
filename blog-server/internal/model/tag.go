package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// gorm omitempty: 如果该字段为空值，那么 gorm 每次将自动忽略该字段

type Tag struct {
	Model

	Name string `gorm:"unique;type:varchar(20);not null" json:"name"`

	// 反向关联该标签下的文章
	Articles []*Article `gorm:"many2many:article_tag" json:"articles,omitempty"`
}

type TagVO struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name         string `json:"name"`
	ArticleCount int    `json:"article_count"`
}

func CreateOrUpdateTag(db *gorm.DB, id int, name string) (*Tag, error) {
	tag := Tag{
		Model: Model{ID: id},
		Name:  name,
	}

	var result *gorm.DB

	if id > 0 {
		result = db.Updates(&tag)
	} else {
		result = db.Create(&tag)
	}

	return &tag, result.Error
}

func GetTagList(db *gorm.DB, pageNum, pageSize int, keyword string) (list []TagVO, total int64, err error) {
	db = db.Table("tag t").
		Joins("left join article_tag at on at.tag_id = t.id").
		Select("t.id", "t.name", "count(at.article_id) as article_count", "t.created_at", "t.updated_at")

	if keyword != "" {
		db = db.Where("name like ?", "%"+keyword+"%")
	}

	err = db.Group("t.id").Order("t.updated_at desc").Count(&total).
		Scopes(Paginate(pageNum, pageSize)).Find(&list).Error

	return list, total, err
}

func DeleteTag(db *gorm.DB, id int) error {
	var articleCount int64
	err := db.Model(&ArticleTag{}).Where("tag_id = ?", id).Count(&articleCount).Error

	if err != nil {
		return err
	}

	if articleCount > 0 {
		return errors.New("tag has articles")
	}

	result := db.Where("id = ?", id).Delete(&Tag{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}