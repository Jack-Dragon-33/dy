package douyin_core

import (
	"context"
	cache "dy/biz/cache"
	dao "dy/biz/db"
	douyin_core "dy/biz/model/douyin_core"
	"dy/biz/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ReleaseComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinCommentActionRequest
	err = c.BindAndValidate(&req)
	resp := new(douyin_core.DouyinCommentActionResponse)
	if err != nil {
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	t := req.ActionType
	mc, err := util.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	var res int32
	var comm *douyin_core.Comment
	if t == 1 {
		comm, res = dao.ReleassComment(mc.UserId, req.VideoId, req.CommentText)
		if res == 1 {
			cache.IncrByVideosComment(req.VideoId)
		}
	} else {
		res = dao.DeleteComment(req.CommentId)
		if res == 1 {
			cache.DecrByVideosComment(req.VideoId)
		}
	}
	if res != 1 {
		resp.StatusCode = FAILD
		resp.StatusMsg = GetStatusMsg(FAILD)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	FillCommentUser(comm)
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	resp.Comment = comm
	c.JSON(consts.StatusOK, resp)
}
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinCommentListRequest
	err = c.BindAndValidate(&req)
	resp := new(douyin_core.DouyinCommentListResponse)
	if err != nil {
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	cs := dao.CommonList(req.VideoId)
	for _, c := range cs {
		FillCommentUser(c)
	}
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	resp.CommentList = cs
	c.JSON(consts.StatusOK, resp)
}
