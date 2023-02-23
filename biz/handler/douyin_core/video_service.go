package douyin_core

import (
	"context"
	dao "dy/biz/db"
	douyin_core "dy/biz/model/douyin_core"
	"dy/biz/util"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	hlog.Info(fmt.Printf("next_time:%d", *req.LatestTime))
	resp := new(douyin_core.DouyinFeedResponse)
	if err != nil {
		resp.StatusCode = ERR_PARMER
		resp.StatusMsg = GetStatusMsg(ERR_PARMER)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	time := req.LatestTime
	hlog.Info("time: %v\n", time)
	videoList := dao.SelectVideoList(time)
	hlog.Info(videoList)
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	var user_id int64
	if req.Token != nil {
		mc, err2 := util.ParseToken(*req.Token)
		if err2 != nil {
			user_id = mc.UserId
		}
	}

	if len(videoList) > 0 {
		for _, v := range videoList {
			FillVideo(v, user_id)
		}
		i := videoList[0].UpdatedAt.Unix()
		resp.NextTime = &i
		resp.VideoList = videoList
		c.JSON(consts.StatusOK, resp)
	}
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinPublishListRequest
	err = c.BindAndValidate(&req)
	resp := new(douyin_core.DouyinPublishListResponse)
	if err != nil {
		resp.StatusCode = ERR_PARMER
		resp.StatusMsg = GetStatusMsg(ERR_PARMER)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	i := req.UserId
	vl := dao.SelectVideoListAll(i)
	var user_id int64
	if len(req.Token)>0 {
		mc, err2 := util.ParseToken(req.Token)
		if err2 != nil {
			user_id = mc.UserId
		}
	}
	for _, v := range vl {
		FillVideo(v,user_id)
	}
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	resp.VideoList = vl
	c.JSON(consts.StatusOK, resp)
}
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin_core.DouyinPublishActionRequest
	c.BindAndValidate(&req)
	fh, err := c.FormFile("data")
	if fh != nil {
		req.Data = fh
	}
	fmt.Printf("fh.Filename: %v\n", fh.Filename)
	hlog.Info(fmt.Printf("req: %v", req))
	resp := new(douyin_core.DouyinPublishActionResponse)
	if err != nil {
		resp.StatusCode = ERR_PARMER
		resp.StatusMsg = GetStatusMsg(ERR_PARMER)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	hlog.Info(fmt.Printf("video %v", req))
	err = c.SaveUploadedFile(req.Data, fmt.Sprintf("biz/static/%s", req.Data.Filename))
	if err != nil {
		hlog.Error(err)
		resp.StatusCode = UPLODA_ERROR
		resp.StatusMsg = GetStatusMsg(UPLODA_ERROR)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	s := util.GetURL(fh.Filename)
	mc, err := util.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = UPLODA_ERROR
		resp.StatusMsg = GetStatusMsg(UPLODA_ERROR)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	hlog.Info(fmt.Printf("mc:%v", mc))
	flag := dao.SaveVideo(s, mc.UserId, req.Title)
	if flag {
		resp.StatusCode = UPLODA_ERROR
		resp.StatusMsg = GetStatusMsg(UPLODA_ERROR)
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}
	resp.StatusCode = SUCCESS
	resp.StatusMsg = GetStatusMsg(SUCCESS)
	c.JSON(consts.StatusOK, resp)
}
