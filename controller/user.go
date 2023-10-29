package controller

import (
	"errors"

	"github.com/Cheng1622/web-short-video/model"
	"github.com/Cheng1622/web-short-video/orm"
	"github.com/Cheng1622/web-short-video/pkg/app"
	"github.com/Cheng1622/web-short-video/pkg/errcode"
	"github.com/Cheng1622/web-short-video/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	a := new(model.User)
	err := orm.DB.Where("Email = ?", c.GetString("Email")).First(&a).Error
	if err != nil {
		slog.Info("GetUser:", err)
		app.ResponseError(c, errcode.CodeTokenInvalid)
		return
	}
	app.ResponseSuccess(c, a)
}

func RegisterUser(c *gin.Context) {
	a := new(model.User)
	if err := c.ShouldBindJSON(a); err != nil {
		slog.Info("RegisterUser invalid param:", err)
		app.ResponseError(c, errcode.CodeInvalidParam)
		return
	}
	err := orm.DB.Create(a).Error
	if err != nil {
		slog.Info("register failed", a.Email, err)
		app.ResponseError(c, errcode.CodeUserExist)
		return
	}
	app.ResponseSuccess(c, a)
}

func LoginUser(c *gin.Context) {
	a := new(model.User)
	if err := c.ShouldBindJSON(a); err != nil {
		slog.Info("LoginUser invalid param:", err)
		app.ResponseError(c, errcode.CodeInvalidParam)
		return
	}

	err := orm.DB.Where("Password = ? AND Email = ?", a.Password, a.Email).First(&a).Error
	if err != nil {
		slog.Info("login failed", a.Email, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.ResponseError(c, errcode.CodeInvalidPassword)
		} else {
			app.ResponseError(c, errcode.CodeServerBusy)
		}
		return
	}
	token, err := util.GenToken(a.Email, a.ID)
	if err != nil {
		slog.Info("token failed", err)
		app.ResponseError(c, errcode.CodeServerBusy)
		return
	}
	app.ResponseSuccess(c, token)
}

func AlterUser(c *gin.Context) {
	a := new(model.User)

	err := orm.DB.Where("ID = ?", c.GetUint("UserId")).Updates(&a).Error
	if err != nil {
		slog.Info("AlterUser failed", c.GetUint("UserId"), err)
		app.ResponseError(c, errcode.CodeUpdatesFail)
		return
	}
	app.ResponseSuccess(c, errcode.CodeSuccess)
}

func DelUser(c *gin.Context) {
	a := new(model.User)
	if err := c.ShouldBindJSON(a); err != nil {
		slog.Info("RegisterUser invalid param:", err)
		app.ResponseError(c, errcode.CodeInvalidParam)
		return
	}
	err := orm.DB.Delete(&a, c.GetUint("UserId"))
	if err != nil {
		slog.Info("DelUser failed", a.Email, err)
		app.ResponseError(c, errcode.CodeDeleteFail)
		return
	}
	app.ResponseSuccess(c, errcode.CodeSuccess)
}
