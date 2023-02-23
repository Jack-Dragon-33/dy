package douyin_core

import "mime/multipart"

type DouyinUserRequest struct {
	UserId int64  `json:"user_id,omitempty" form:"user_id" query:"user_id"` // 用户id
	Token  string `json:"token,omitempty" form:"token" query:"token"`       // 用户鉴权token
}
type DouyinUserResponse struct {
	StatusCode int32   `json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`    // 返回状态描述
	User       *User   `json:"user,omitempty" form:"user" query:"user"`                      // 用户信息
}
type DouyinUserLoginRequest struct {
	Username string `json:"username,omitempty" query:"username"` // 登录用户名
	Password string `json:"password,omitempty" query:"password"` // 登录密码
}
type DouyinUserLoginResponse struct {
	StatusCode int32   `json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`    // 返回状态描述
	UserId     int64   `json:"user_id,omitempty" form:"user_id" query:"user_id"`             // 用户id
	Token      string  `json:"token,omitempty" form:"token" query:"token"`                   // 用户鉴权token
}
type DouyinUserRegisterRequest struct {
	Username string `json:"username,omitempty" form:"username" query:"username"` // 注册用户名，最长32个字符
	Password string `json:"password,omitempty" form:"password" query:"password"` // 密码，最长32个字符
	Name     string `json:"name" query:"name" form:"name"`
}
type DouyinUserRegisterResponse struct {
	StatusCode int32   `json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`    // 返回状态描述
	UserId     int64   `json:"user_id,omitempty" form:"user_id" query:"user_id"`             // 用户id
	Token      string  `json:"token,omitempty" form:"token" query:"token"`                   // 用户鉴权token
}
type DouyinFeedRequest struct {
	LatestTime *int64  `json:"latest_time,omitempty" form:"latest_time" query:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty" form:"token" query:"token"`                   // 可选参数，登录用户设置
}
type DouyinFeedResponse struct {
	StatusCode int32    `json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`    // 返回状态描述
	VideoList  []*Video `json:"video_list" form:"video_list" query:"video_list"`              // 视频列表
	NextTime   *int64   `json:"next_time,omitempty" form:"next_time" query:"next_time"`       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
type DouyinPublishListRequest struct {
	UserId int64  `json:"user_id" form:"user_id" query:"user_id"`
	Token  string `json:"token" form:"token" query:"token"`
}
type DouyinPublishListResponse struct {
	StatusCode int32    `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string  `json:"status_msg " form:"status_msg" query:"status_msg"`
	VideoList  []*Video `json:"video_List" form:"status_msg" query:"status_msg"`
}
type DouyinPublishActionRequest struct {
	Token string `json:"token" form:"token" query:"token"`
	Title string `json:"title" form:"title" query:"title"`
	Data  *multipart.FileHeader
}
type DouyinPublishActionResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string `json:"status_msg " form:"status_msg" query:"status_msg"`
}
type DouyinFavoriteActionRequest struct {
	Token      string `json:"token" form:"token" query:"token"`
	VideoId    int64  `json:"video_id" form:"video_id" query:"video_id"`
	ActionType int32  `json:"action_type" form:"action_type" query:"action_type"`
}
type DouyinFavoriteActionResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string `json:"status_msg " form:"status_msg" query:"status_msg"`
}
type DouyinCommentActionRequest struct {
	Token       string `json:"token" form:"token" query:"token"`
	VideoId     int64  `json:"video_id" form:"video_id" query:"video_id"`
	ActionType  int32  `json:"action_type" form:"action_type" query:"action_type"`
	CommentText string `json:"comment_text" form:"comment_text" query:"comment_text"`
	CommentId   int64  `json:"comment_id" form:"comment_id" query:"comment_id"`
}
type DouyinCommentActionResponse struct {
	StatusCode int32    `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string  `json:"status_msg " form:"status_msg" query:"status_msg"`
	Comment    *Comment `json:"comment"`
}
type DouyinFavoriteListRequest struct{
	UserId int64 `json:"user_id" form:"user_id" query:"user_id"`
	Token string `json:"token" form:"token" query:"token"`
}
type DouyinFavoriteListResponse struct{
	StatusCode int32    `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string  `json:"status_msg " form:"status_msg" query:"status_msg"`
	VideoList  []*Video `json:"video_List" form:"status_msg" query:"status_msg"`
}
type DouyinCommentListRequest struct{
	Token      string `json:"token" form:"token" query:"token"`
	VideoId    int64  `json:"video_id" form:"video_id" query:"video_id"`
}
type DouyinCommentListResponse struct{
	StatusCode int32    `json:"status_code" form:"status_code" query:"status_code"`
	StatusMsg  *string  `json:"status_msg " form:"status_msg" query:"status_msg"`
	CommentList []*Comment `json:"comment_list" form:"comment_list" query:"comment_list"`
}