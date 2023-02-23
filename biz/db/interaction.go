package db

import (
	douyin_core "dy/biz/model/douyin_core"
)

func FavoriteAction(userId int64, videoId int64) (int32, error) {
	favoite_modal := douyin_core.FavoriteVideo{
		UserID:  userId,
		VideoId: videoId,
	}
	d := DB.Debug().Create(&favoite_modal)
	return int32(d.RowsAffected), d.Error
}
func CacalFavoritAction(userId int64, videoId int64) int32 {
	favoite_modal := douyin_core.FavoriteVideo{
		UserID:  userId,
		VideoId: videoId,
	}
	d := DB.Debug().Where("video_id =? and user_id=?", videoId, userId).Delete(&favoite_modal)
	return int32(d.RowsAffected)
}
func GetUserFavoriteList() []*douyin_core.UserFavoriteVideo {
	var ufv []*douyin_core.UserFavoriteVideo
	DB.Debug().Select("video_id", "user_id").Find(&ufv)
	return ufv
}
func GetFavoriteVideos(user_id int64) []*douyin_core.Video {
	var vs []*douyin_core.Video
	DB.Debug().Order("updated_at asc").Where("id IN(SELECT video_id FROM favorit_videos WHERE user_id =?)", user_id).Limit(VIDEO_LIMIT).Preload("Author").Find(&vs)
	return vs
}
