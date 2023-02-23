package cache

import (
	dao "dy/biz/db"
	douyin_core "dy/biz/model/douyin_core"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func InitUserCount() {
	p := RedisClient.Pipeline()
	vc := GetUserCount()
	for _, v := range vc {
		key := GetCountKey(USER, v.UserId)
		rw, err := p.Del(ctx, key).Result()
		if err != nil {
			hlog.Info(ctx, "del %s, rw=%d\n", key, rw)
		}
		_, err = p.HMSet(ctx, key, v).Result()
		if err != nil {
			hlog.Error(err)
		}
		hlog.Info(fmt.Printf("设置 user_id=%d, key=%s\n", v.UserId, key))
	}
	_, err := p.Exec(ctx)
	if err != nil {
		_, err = p.Exec(ctx)
		if err != nil {
			hlog.Error(err)
		}
	}

}
func InitUserFavorite() {
	p := RedisClient.Pipeline()
	ufv := GetUserFavoriteList()
	for _, v := range ufv {
		key := GetCountKey(USER_FAVORITE, v.UserId)
		rw, err := p.Del(ctx, key).Result()
		if err != nil {
			hlog.Info(ctx, "del %s, rw=%d\n", key, rw)
		}
		_, err = p.SAdd(ctx, key, v.VideoId).Result()
		if err != nil {
			hlog.Error(err)
		}
		hlog.Info(fmt.Printf("设置 user_id=%d, key=%s\n", v.UserId, key))
	}
	_, err := p.Exec(ctx)
	if err != nil {
		_, err = p.Exec(ctx)
		if err != nil {
			hlog.Error(err)
		}
	}
}
func GetUserCount() []*douyin_core.UserCount {
	return dao.UserCount()
}

func IncrByUserFavorite(user_id int64) {
	IncrByFiled(ctx, USER, user_id, "favorite_count")
}
func DecrByUserFavorite(user_id int64) {
	DecrByFiled(ctx, USER, user_id, "favorite_count")
}
func IncrByUserTotalFavorite(user_id int64) {
	IncrByFiled(ctx, USER, user_id, "total_favorited")
}
func DecrByUserTotalFavorite(user_id int64) {
	DecrByFiled(ctx, USER, user_id, "total_favorited")
}
func IncrByUserWorkCount(user_id int64) {
	IncrByFiled(ctx, USER, user_id, "work_count")
}
func DecrByUserWorkCount(user_id int64) {
	DecrByFiled(ctx, USER, user_id, "work_count")
}

func GetUserCountByUserId(user_id int64) *douyin_core.UserCount {
	var usercount douyin_core.UserCount
	err := RedisClient.HGetAll(ctx, GetCountKey(USER, user_id)).Scan(&usercount)
	if err != nil {
		hlog.Error(err)
	}
	return &usercount
}

func GetUserFavoriteList() []*douyin_core.UserFavoriteVideo {
	return dao.GetUserFavoriteList()
}

func IsUserFavorite(user_id int64, video_id int64) bool {
	key := GetCountKey(USER_FAVORITE, user_id)
	bc := RedisClient.SIsMember(ctx, key, video_id)
	return bc.Val()
}
