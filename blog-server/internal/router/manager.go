package router

import (
	"gin-blog/internal/handle"
	"github.com/gin-gonic/gin"
)

// handle instance
var (
	blogViewApi handle.BlogInfo // blog 前台展示 handle 处理函数
	articleApi  handle.Article  // 文章相关 handle 处理函数
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
	blog := r.Group("/api/v1/blog")

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
		article.GET("/list", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List all articles",
			})
		})
		article.POST("/create", articleApi.CreateArticle)
		article.POST("/update", articleApi.UpdateArticle)
		article.GET("/:id", articleApi.GetArticle)
		article.DELETE("/:id", articleApi.DeleteArticle)
		article.PUT("/soft-delete/:id", articleApi.SoftDeleteArticle)
	}

}
