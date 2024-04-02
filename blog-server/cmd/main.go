package main

import (
	router "gin-blog/blog-server/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	router.RegisterRouter(r)
}
