package db

import (
	douyin_core "dy/biz/model/douyin_core"
	utils "dy/biz/util"
)

func UserRegister(user *douyin_core.User) int64 {
	id := utils.GenerateID()
	user.Id = id
	d := DB.Select("id", "name", "username", "password", "created_at", "updated_at", "deleted_at").Create(&user)
	if d.RowsAffected == 0 {
		return -1
	}
	return user.Id
}

func UserInfo(username string) []*douyin_core.User {
	var users []*douyin_core.User
	DB.Where("username=?", username).Find(&users)
	return users
}

func UserCount() []*douyin_core.UserCount {
	var UserCounts []*douyin_core.UserCount
	DB.Debug().Raw("SELECT u.id user_id,IFNULL(t1.favorite_count,0) favorite_count,IFNULL(t1.total_favorited,0) total_favorited,IFNULL(t1.work_count,0) work_count FROM users u LEFT JOIN(SELECT user_id ,SUM(t.favorite_count) favorite_count,SUM(t.total_favorited)total_favorited,SUM(t.work_count) work_count FROM	(SELECT user_id,COUNT(*) favorite_count ,0 total_favorited, 0 work_count FROM favorit_videos GROUP BY user_id UNION ALL	SELECT v.user_id ,0 favorite_count, 0 work_count, COUNT(*) total_favorited FROM videos v LEFT JOIN favorit_videos f ON v.id =f.video_id GROUP BY v.user_id	UNION ALL SELECT user_id,0 favorite_count,0 total_favorited, COUNT(*) workcount FROM videos GROUP BY user_id) t GROUP BY user_id) t1 ON u.id=t1.user_id").Scan(&UserCounts)
	return UserCounts
}
