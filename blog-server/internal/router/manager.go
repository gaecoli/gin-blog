package router

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	registerJwtHandler(r)
	registerBlogViewHandler(r)
	registerBlogManagerHandler(r)
}

func registerJwtHandler(r *gin.Engine) {

}

func registerBlogManagerHandler(r *gin.Engine) {

}

//  blog view handler, don't need to auth
func registerBlogViewHandler(r *gin.Engine) {
	base := r.Group("/api/v1")

	base.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the blog!",
		})
	})

	article := base.Group("/article")
	{

		article.GET("/list", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List all articles",
			})
		})
	}

}
