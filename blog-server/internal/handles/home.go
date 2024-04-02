package handles

import "github.com/gin-gonic/gin"

type BlogView struct{}

func (*BlogView) BlogHome(c *gin.Context) {
	c.String(200, "Welcome to my blog!")
}
