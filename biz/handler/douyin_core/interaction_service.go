package douyin_core

import (
	"context"
	dao "dy/biz/db"
	douyin_core "dy/biz/model/douyin_core"
	"dy/biz/util"
	"fmt"

	cache "dy/biz/cache"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinFavoriteActionRequest
	err = c.BindAndValidate(&req)
	hlog.Info(fmt.Printf("token:%s video_id:%d", req.Token, req.VideoId))
	resp := new(douyin_core.DouyinFavoriteActionResponse)
	if err != nil {
		hlog.Error(err)
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	t := req.ActionType
	mc, err := util.ParseToken(req.Token)
	hlog.Info(fmt.Printf("mc: %v", mc))
	if err != nil {
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	var res int32
	if t == 1 {
		res, err = dao.FavoriteAction(mc.UserId, req.VideoId)
		if err != nil {
			resp.StatusCode = REPEAT_ACTION
			resp.StatusMsg = GetStatusMsg(REPEAT_ACTION)
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
	} else {
		res = dao.CacalFavoritAction(mc.UserId, req.VideoId)
	}
	if res != 1 {
		resp.StatusCode = FAILD
		resp.StatusMsg = GetStatusMsg(FAILD)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	cache.FavoriteAction(req.VideoId, mc.UserId, req.ActionType)
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	c.JSON(consts.StatusOK, resp)
}
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinFavoriteListRequest
	resp := new(douyin_core.DouyinFavoriteListResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		hlog.Error(err)
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
		return
	}
	var mc *douyin_core.MyClaim
	mc, err = util.ParseToken(req.Token)
	if err!=nil {
		resp.StatusCode = LOGIN_NO_REQUEST
		resp.StatusMsg = GetStatusMsg(LOGIN_NO_REQUEST)
		c.JSON(consts.StatusOK, resp)
	}
	vs := dao.GetFavoriteVideos(req.UserId)
	for _, v := range vs {
		FillVideo(v, mc.UserId)
	}
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	resp.VideoList = vs
	c.JSON(consts.StatusOK, resp)
}
