package main

import (
	"flag"
	g "gin-blog/internal/global"
	middle "gin-blog/internal/middleware"
	"gin-blog/internal/model"
	"gin-blog/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configPath := flag.String("c", "./config.yml", "配置文件路径")
	flag.Parse()

	conf := g.ReadConfig(*configPath)

	_ = middle.InitLogger(conf)

	db := model.InitDB(conf)

	// set gin router
	gin.SetMode(conf.Server.Mode)
	r := gin.New()
	err := r.SetTrustedProxies([]string{"*"})
	if err != nil {
		return
	}

	if conf.Server.Mode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(middle.Recovery(true), middle.GinLogger())
	}
	// add gin cors middleware
	r.Use(middle.Cors())
	// gin handle with gorm
	r.Use(middle.WithGormDB(db))

	router.RegisterRouter(r)

	serverPort := conf.Server.Port

	log.Printf("gin-blot serveing HTTP on (http://%s/)...\n", serverPort)

	r.Run(serverPort)
}
