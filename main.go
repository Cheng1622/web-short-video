package main

import (
	"github.com/Cheng1622/web-short-video/config"
	"github.com/Cheng1622/web-short-video/controller"
	"github.com/Cheng1622/web-short-video/model"
	"github.com/Cheng1622/web-short-video/orm"
	"github.com/gin-gonic/gin"
)

func main() {
	orm.InitSqlite()
	orm.DB.AutoMigrate(&model.Video{})

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", controller.GetVideo)
	}

	router.Run(config.PORT)
}
