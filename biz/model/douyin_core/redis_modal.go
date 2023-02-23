package douyin_core

type VideoCount struct {
	VideoId       int64 `json:"video_id"`
	FavoriteCount int64 `redis:"favorite_count" json:"favorite_count"`
	CommentCount  int64 `redis:"comment_count" json:"comment_count"`
}
type UserCount struct {
	UserId         int64 `json:"user_id"`
	TotalFavorited int64 `redis:"total_favorited" json:"total_favorited"`
	WorkCount      int64 `redis:"work_count" json:"work_count"`
	FavoriteCount  int64 `redis:"favorite_count" json:"favorite_count"`
}
type UserFavoriteVideo struct{
	VideoId int64 `json:"video_id"`
	UserId int64 `json:"user_id"`
}
func (UserFavoriteVideo) TableName() string {
	return "favorit_videos"
}
