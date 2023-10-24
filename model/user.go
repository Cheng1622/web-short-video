package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name          string `json:"name"`           // 用户名称
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}
