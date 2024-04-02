package router

import (
	"gin-blog/blog-server/internal/handles"
	"github.com/gin-gonic/gin"
)

var (
	blogViewApi handles.BlogView
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

	blog.GET("/home", blogViewApi.BlogHome)
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
	}

}
