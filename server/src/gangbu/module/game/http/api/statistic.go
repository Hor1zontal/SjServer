package api

import (
	"gangbu/module/game/http/helper"
	"gangbu/module/game/service/statistic"
	"github.com/gin-gonic/gin"
)

type NewlyReq struct {
	Reg   string `form:"reg"`
	Active string `form:"active"`
}

type NewlyResp struct {
	RegCount int32 `json:"reg_count"`
	ActiveCount int32 `json:"active_count"`
}

//新增留存
func NewlyRemain(c *gin.Context) {
	req := &NewlyReq{}
	helper.CheckReq(c, req)
	reg_count, active_count := statistic.GetStatisticCount("regtime", "activeTime", req.Reg, req.Active)
	resp := &NewlyResp{RegCount:reg_count, ActiveCount:active_count}
	helper.ResponseWithData(c, resp)
}

type ActiveReq struct {
	LastActive string `form:"last_active"`
	Active string `form:"active"`
}

type ActiveResp struct {
	LastCount int32 `json:"last_active"`
	Count int32 `json:"count"`
}

//活跃留存
func ActivityRemain(c *gin.Context) {
	req := &ActiveReq{}
	helper.CheckReq(c, req)
	last_count, count := statistic.GetStatisticCount("lastActiveTime", "activeTime", req.LastActive, req.Active)
	resp := &ActiveResp{LastCount:last_count, Count:count}
	helper.ResponseWithData(c, resp)
}

