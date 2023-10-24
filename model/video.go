package model

import (
	"gorm.io/gorm"
)

type Video struct {
	*gorm.Model
	Title         string `gorm:"not null"`                     // 视频标题
	Author        User   `gorm:"foreignKey:AuthorID;not null"` // 视频作者信息
	CoverURL      string `gorm:"not null"`                     // 视频封面地址
	PlayURL       string `gorm:"not null"`                     // 视频播放地址
	FavoriteCount int64  `gorm:"not null;default:0"`           // 视频的点赞总数
	CommentCount  int64  `gorm:"not null;default:0"`           // 视频的评论总数
}
