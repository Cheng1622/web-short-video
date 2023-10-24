package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name          string `gorm:"not null"`           // 用户名称
	FollowCount   int64  `gorm:"not null;default:0"` // 关注总数
	FollowerCount int64  `gorm:"not null;default:0"` // 粉丝总数
}
