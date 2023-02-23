package cache

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "192.168.200.130:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	RedisClient = rdb
	initData()
	hlog.Info("redis 初始化完成")
}
func initData() {
	InitVideoCount()
	InitUserCount()
	InitUserFavorite()
}
