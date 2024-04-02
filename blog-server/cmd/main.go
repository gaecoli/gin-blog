package main

import (
	middle "gin-blog/blog-server/internal/middleware"
	"gin-blog/blog-server/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.SetTrustedProxies([]string{"*"})
	// add gin cors middleware
	r.Use(middle.Cors())

	router.RegisterRouter(r)

	r.Run(":9099")
}
