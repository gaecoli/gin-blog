package router

import (
	"gin-blog/internal/handle"
	"github.com/gin-gonic/gin"
)

// handle instance
var (
	blogViewApi handle.BlogInfo // blog 前台展示 handle 处理函数
	articleApi  handle.Article  // 文章相关 handle 处理函数
	categoryApi handle.Category // 分类相关 handle 处理函数
	tagApi      handle.Tag
)

func RegisterRouter(r *gin.Engine) {
	registerJwtHandler(r)
	registerBlogViewHandler(r)
	registerBlogManagerHandler(r)
}

// jwt handler, include login, register, etc.
func registerJwtHandler(r *gin.Engine) {

}

// blog manager handler, need to jwt auth
func registerBlogManagerHandler(r *gin.Engine) {

}

// blog view handler, don't need to auth
func registerBlogViewHandler(r *gin.Engine) {
	blog := r.Group("/api/v1")

	blog.GET("/home", blogViewApi.GetHomeInfo)
	blog.GET("/page", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Page",
		})
	})
	blog.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "About me",
		})
	})

	article := blog.Group("/article")
	{
		article.GET("/list", articleApi.GetArticleList)
		article.POST("/create", articleApi.CreateArticle)
		article.POST("/update", articleApi.UpdateArticle)
		article.GET("/:id", articleApi.GetArticle)
		article.DELETE("/:id", articleApi.DeleteArticle)
		article.PUT("/archive", articleApi.SoftDeleteArticle)
	}

	category := blog.Group("/category")
	{
		category.GET("/list", categoryApi.GetCategoryList)
		category.POST("/create", categoryApi.CreateCategory)
		category.POST("/update", categoryApi.UpdateCategory)
		category.DELETE("/:id", categoryApi.DeleteCategory)
	}

	tag := blog.Group("/tag")
	{
		tag.POST("", tagApi.CreateOrUpdateTag)
		tag.GET("/list", tagApi.GetTagList)
		tag.DELETE("/:id", tagApi.DeleteTag)
	}
}
