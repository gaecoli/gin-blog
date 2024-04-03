package main

import (
	"fmt"
	middle "gin-blog/blog-server/internal/middleware"
	"gin-blog/blog-server/internal/model"
	"gin-blog/blog-server/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// init db
	err := model.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	// close db
	defer func() {
		err := model.CloseDB()
		if err != nil {
			fmt.Println(err)
		}
	}()

	// set gin router
	r := gin.New()
	r.SetTrustedProxies([]string{"*"})
	// add gin cors middleware
	r.Use(middle.Cors())

	router.RegisterRouter(r)

	r.Run(":9099")
}
