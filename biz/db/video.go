package db

import (
	douyin_core "dy/biz/model/douyin_core"
	utils "dy/biz/util"
)

const (
	VIDEO_LIMIT = 5
)

func SelectVideoList(curtime *int64) []*douyin_core.Video {
	var videos []*douyin_core.Video
	t := utils.ParseTime(*curtime)
	DB.Debug().Order("updated_at asc").Where("updated_at>?", t).Limit(VIDEO_LIMIT).Preload("Author").Find(&videos)
	return videos
}
func SelectVideoListAll(userId int64) []*douyin_core.Video {
	var videos []*douyin_core.Video
	DB.Debug().Order("updated_at asc").Where("user_id=?", userId).Limit(VIDEO_LIMIT).Preload("Author").Find(&videos)
	return videos
}
func SaveVideo(filePath string, userId int64, title string) bool {
	id := utils.GenerateID()
	v := douyin_core.Video{Id: id, UserId: userId, Title: title, PlayUrl: filePath, FavoriteCount: 0, CommentCount: 0}
	d := DB.Debug().Select("id", "user_id", "play_url", "cover_url", "title", "created_at", "updated_at", "deleted_at").Create(&v)
	return d.RowsAffected > 0
}
func GetVideoCount() []*douyin_core.VideoCount {
	var vf []*douyin_core.VideoCount
	DB.Debug().Raw("SELECT v.id video_id ,IFNULL(t1.favorite_count,0) favorite_count,IFNULL(t1.comment_count,0) comment_count FROM videos v LEFT JOIN(SELECT t.video_id video_id,SUM(t.favorite_count) favorite_count,SUM(t.comment_count) comment_count FROM (	SELECT video_id,COUNT(*) favorite_count ,0 comment_count FROM favorit_videos f GROUP BY video_id UNION SELECT video_id,0 favorite_count ,COUNT(*) comment_count FROM comment_videos GROUP BY video_id) t GROUP BY video_id)t1 ON t1.video_id=v.id").Scan(&vf)
	return vf
}
