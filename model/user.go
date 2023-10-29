package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	UserName      string `gorm:"type:varchar(50);not null"`
	Password      string `gorm:"type:varchar(64);not null"`
	Email         string `gorm:"type:varchar(150);not null;unique"`
	ImageUrl      string `gorm:"type:varchar(255);not null"`
	Status        int64  `gorm:"not null;default:0"`
	FollowCount   int64  `gorm:"not null;default:0"` // 关注总数
	FollowerCount int64  `gorm:"not null;default:0"` // 粉丝总数
}
