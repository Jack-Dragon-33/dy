package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	VIDEO         = 1
	USER          = 2
	USER_FAVORITE = 3
)

var (
	prefixMap = map[int16]string{
		VIDEO:         "dy_video_count",
		USER:          "dy_user_count",
		USER_FAVORITE: "dy_user_favorite",
	}
	ctx = context.Background()
)

func DecrByFiled(ctx context.Context, key_type int16, id int64, field string) {
	change(ctx, key_type, id, field, -1)
}

func IncrByFiled(ctx context.Context, key_type int16, id int64, field string) {
	change(ctx, key_type, id, field, 1)
}

func change(ctx context.Context, key_type int16, ID int64, field string, incr int64) {
	redisKey := GetCountKey(key_type, ID)
	before, err := RedisClient.HGet(ctx, redisKey, field).Result()
	if err != nil {
		panic(err)
	}
	beforeInt, err := strconv.ParseInt(before, 10, 64)
	if err != nil {
		panic(err)
	}
	if beforeInt+incr < 1 {
		hlog.Info(fmt.Printf("禁止变更计数,计数变更后小于0. %d + (%d) = %d\n", beforeInt, incr, beforeInt+incr))
		return
	}
	hlog.Info(fmt.Printf("Id: %d\n更新前\n%s = %s\n--------\n", ID, field, before))
	_, err = RedisClient.HIncrBy(ctx, redisKey, field, incr).Result()
	if err != nil {
		panic(err)
	}
	// fmt.Printf("更新记录[%d]:%d\n", userID, num)
	count, err := RedisClient.HGet(ctx, redisKey, field).Result()
	if err != nil {
		panic(err)
	}
	hlog.Info(fmt.Printf("ID: %d\n更新后\n%s = %s\n--------\n", ID, field, count))
}
func GetCountKey(key_type int16, id int64) string {
	prefix := prefixMap[int16(key_type)]
	return fmt.Sprintf("%s_%d", prefix, id)
}
