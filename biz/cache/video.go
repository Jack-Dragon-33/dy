package cache

import (
	dao "dy/biz/db"
	"dy/biz/model/douyin_core"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitVideoCount() {
	p := RedisClient.Pipeline()
	vc := GetVideoCommentCount()
	for _, v := range vc {
		key := GetCountKey(VIDEO, v.VideoId)
		rw, err := p.Del(ctx, key).Result()
		if err != nil {
			hlog.Info(ctx, "del %s, rw=%d\n", key, rw)
		}
		_, err = p.HMSet(ctx, key, v).Result()
		if err != nil {
			hlog.Error(err)
		}
		hlog.Info(fmt.Printf("设置 video_id=%d, key=%s\n", v.VideoId, key))
	}
	_, err := p.Exec(ctx)
	if err != nil {
		_, err = p.Exec(ctx)
		if err != nil {
			hlog.Error(err)
		}
	}

}
func GetVideoCommentCount() []*douyin_core.VideoCount {
	return dao.GetVideoCount()
}

func IncrByVideosFavorite(video_id int64) {
	IncrByFiled(ctx, VIDEO, video_id, "favorite_count")
}
func DecrByVideosFavorite(video_id int64) {
	DecrByFiled(ctx, VIDEO, video_id, "favorite_count")
}
func IncrByVideosComment(video_id int64) {
	IncrByFiled(ctx, VIDEO, video_id, "comment_count")
}
func DecrByVideosComment(video_id int64) {
	DecrByFiled(ctx, VIDEO, video_id, "comment_count")
}
func GetVideoCountByVideoId(video_id int64) *douyin_core.VideoCount {
	var video_count douyin_core.VideoCount
	RedisClient.HGetAll(ctx, GetCountKey(VIDEO, video_id)).Scan(&video_count)
	return &video_count
}
func AddUserFavorite(user_id int64,video_id int64)(err error){
	key := GetCountKey(USER_FAVORITE, user_id)
	_, err = RedisClient.SAdd(ctx, key, video_id).Result()
	return err
}
func SubUserFavorite(user_id int64,video_id int64)(err error){
	key := GetCountKey(USER_FAVORITE, user_id)
	_, err = RedisClient.SRem(ctx, key, video_id).Result()
	return err
}
func FavoriteAction(video_id int64 ,user_id int64, action_type int32){
	var err error
	switch(action_type){
	case 1:
		IncrByVideosFavorite(video_id)
		IncrByUserFavorite(user_id)
		err = AddUserFavorite(user_id, video_id)
	case 2:
		DecrByVideosFavorite(video_id)
		DecrByUserFavorite(user_id)
		err =SubUserFavorite(user_id,video_id)
	}
	if err!=nil{
		hlog.Error(err)
		return
	}
}
