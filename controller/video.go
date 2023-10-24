package controller

import (
	"github.com/Cheng1622/web-short-video/model"
	"github.com/Cheng1622/web-short-video/orm"
	"github.com/Cheng1622/web-short-video/pkg"
	"github.com/Cheng1622/web-short-video/pkg/app"
	"github.com/gin-gonic/gin"
)

func GetVideo(c *gin.Context) {

	app.ResponseSuccess(c, pkg.New(c, orm.DB.Model(&model.Video{})))
}
