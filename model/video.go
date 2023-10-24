package model

import "gorm.io/gorm"

type Video struct {
	*gorm.Model
	Title         string `json:"title"`          // 视频标题
	Author        User   `json:"author"`         // 视频作者信息
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	PlayURL       string `json:"play_url"`       // 视频播放地址
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
}

type ApifoxModel struct {
	NextTime   *int64  `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 视频列表
}
