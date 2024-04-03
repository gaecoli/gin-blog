package handle

import "github.com/gin-gonic/gin"

// BlogView 前台展示的博客的视图
type BlogView struct{}

func (*BlogView) BlogHome(c *gin.Context) {
	c.String(200, "Welcome to my blog!")
}
