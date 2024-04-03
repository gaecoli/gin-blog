package model

//import "gorm.io/gorm"

// BlogHomeStatistics 返回博客首页统计数据
//type BlogHomeStatistics struct {
//	ArticleCount   int64             `json:"article_count"`
//	UserCount      int64             `json:"user_count"`
//	MessageCount   int64             `json:"message_count"`
//	CategoryCount  int64             `json:"category_count"`
//	TagCount       int64             `json:"tag_count"`
//	ViewCount      int64             `json:"view_count"`
//	BlogConfigInfo map[string]string `json:"blog_config"`
//}
//
//// GetBlogHomeStatistics 获取博客首页统计数据
//func (db *gorm.DB) GetBlogHomeStatistics() (data BlogHomeStatistics, err error) {
//
//}

//type BlogPage struct {
//	Model
//
//	Title       string `json:"title"`
//	Context     string `json:"context"`
//	Description string `json:"description"`
//}
//
//func (db *gorm.DB) GetBlogPage(blogPage BlogPage) (BlogPage, error) {
//	result := db.Where("id = ?", blogPage.ID).First(&blogPage)
//	return blogPage, result.Error
//}
