package douyin_core

import (
	"dy/biz/cache"
	"dy/biz/model/douyin_core"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func fillVideoUser(video *douyin_core.Video) {
	uc := cache.GetUserCountByUserId(video.UserId)
	if uc != nil {
		video.Author.FavoriteCount = uc.FavoriteCount
		video.Author.TotalFavorited = uc.TotalFavorited
		video.Author.WorkCount = uc.WorkCount
	}
}
func fillVideoCount(video *douyin_core.Video, user_id int64) {
	vc := cache.GetVideoCountByVideoId(video.Id)
	b := cache.IsUserFavorite(user_id, video.Id)
	hlog.Info("vc: ", vc)
	if vc != nil {
		video.CommentCount = vc.CommentCount
		video.FavoriteCount = vc.FavoriteCount
		video.IsFavorite = b
		hlog.Info(fmt.Printf("video_id:%d,user_id:%d,is_favorite:%T", video.Id, user_id, b))
	}

}
func FillVideo(video *douyin_core.Video, user_id int64) {
	fillVideoCount(video, user_id)
	fillVideoUser(video)
}
func FillCommentUser(comm *douyin_core.Comment) {
	uc := cache.GetUserCountByUserId(comm.UserID)
	if uc != nil {
		comm.User.TotalFavorited = uc.TotalFavorited
		comm.User.FavoriteCount = uc.FavoriteCount
		comm.User.WorkCount = uc.WorkCount
	}
}
func FillUser(user *douyin_core.User) {
	uc := cache.GetUserCountByUserId(user.Id)
	if uc != nil {
		user.TotalFavorited = uc.TotalFavorited
		user.FavoriteCount = uc.FavoriteCount
		user.WorkCount = uc.WorkCount
	}
}
