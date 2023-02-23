package douyin_core

import (
	_ "dy/biz/model/api"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`        // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"` // 用户名称
	Username        string `json:"username" form:"username" query:"username"`
	Password        string `json:"password" form:"password" query:"password"`
	Avatar          string `json:"avatar" form:"avatar" query:"avatar"`
	BackgroundImage string `json:"background_image" form:"background_image" query:"background_image"`
	Signature       string `json:"signature" form:"signature" query:"signature"`
	TotalFavorited  int64  `json:"total_favorited" form:"total_favorited" query:"total_favorited"`
	WorkCount       int64  `json:"work_count" form:"work_count" query:"work_count"`
	FavoriteCount   int64  `json:"favorite_count" form:"favorite_count" query:"favorite_count"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {

	return nil
}

type MyClaim struct {
	UserId       int64  `json:"userId"`
	Name         string `json:"name"`
	FollowCount  *int64 `json:"follow_count"`
	FollowrCount *int64 `json:"follower_count"`
	IsFollow     bool   `json:"is_follow"`
	jwt.RegisteredClaims
}

type FavoriteVideo struct {
	Id        int64 `json:"id"`
	VideoId   int64 `json:"video_id"`
	UserID    int64 `json:"user_id"`
	CreatedAt time.Time
}

func (FavoriteVideo) TableName() string {
	return "favorit_videos"
}

type Comment struct {
	Id          int64  `json:"id"`
	VideoId     int64  `json:"video_id"`
	UserID      int64  `json:"user_id"`
	User        User   `gorm:"foreignkey:Id;references:user_id" json:"user"`
	CommentText string `json:"comment_text"`
	CreatedAt   time.Time
}

func (Comment) TableName() string {
	return "comment_videos"
}

type Video struct {
	Id            int64  `json:"id,omitempty" form:"id" query:"id"` // 视频唯一标识
	UserId        int64  `json:"user_id"`
	Author        *User  `gorm:"foreignkey:Id;references:user_id" json:"author,omitempty" form:"author" query:"author"` // 视频作者信息
	PlayUrl       string `json:"play_url,omitempty" form:"play_url" query:"play_url"`                                   // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty" form:"cover_url" query:"cover_url"`                                // 视频封面地址
	FavoriteCount int64  `json:"favorite_count" form:"favorite_count" query:"favorite_count"`                 // 视频的点赞总数
	CommentCount  int64  `json:"comment_count" form:"comment_count" query:"comment_count"`                    // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite" form:"is_favorite" query:"is_favorite"`                          // true-已点赞，false-未点赞
	Title         string `json:"title" form:"title" query:"title"`                                                                   // 视频标题
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
