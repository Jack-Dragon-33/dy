package db

import (
	douyin_core "dy/biz/model/douyin_core"
	utils "dy/biz/util"
)

func ReleassComment(userId int64, videoId int64, commentText string) (*douyin_core.Comment, int32) {
	i := utils.GenerateID()
	comm := douyin_core.Comment{
		Id:          i,
		UserID:      userId,
		VideoId:     videoId,
		CommentText: commentText,
	}
	d := DB.Debug().Create(&comm)
	return &comm, int32(d.RowsAffected)
}
func DeleteComment(comment_id int64) int32 {
	comment := douyin_core.Comment{
		Id: comment_id,
	}
	d := DB.Debug().Delete(&comment)
	return int32(d.RowsAffected)
}
func CommonList(video_id int64) []*douyin_core.Comment {
	var comms []*douyin_core.Comment
	DB.Debug().Order("created_at DESC ").Preload("User").Where("video_id =?", video_id).Find(&comms)
	return comms
}
