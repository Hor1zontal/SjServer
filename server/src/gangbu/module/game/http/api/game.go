package api

import (
	"encoding/json"
	"gangbu/exception"
	"gangbu/module/game/http/helper"
	"gangbu/module/game/service"
	"github.com/gin-gonic/gin"
	"github.com/name5566/leaf/log"
)

//type EnterGameReq struct {
//	Props []*db.Prop	`form:"props"`
//}

type EnterGameReq struct {
	Guide	bool `json:"guide"`
}

type EnterGameResp struct {
	Energy	int32	`json:"energy"`
	EnergyTime int64 `json:"energyTime"`
}

func EnterGame(c *gin.Context) {
	req := &EnterGameReq{}
	helper.CheckReq(c, req)
	service.EnterGameData(helper.GetClaimUser(c).UID/*, req.Props*/)
	role := service.UseEnergy(helper.GetClaimUser(c).UID,req.Guide)
	resp := &EnterGameResp{
		Energy:role.Energy,
		EnergyTime:role.EnergyTime.Unix(),
	}
	helper.ResponseWithData(c, resp)
}

type UpdateGameReq struct {
	Guide	bool `form:"guide"`
	Floor	int32 `form:"floor"`
	Score	int32 `form:"score" binding:"required"`
	//Props   []*db.Prop `form:"props"`
	Items	string `form:"items"`
	//BoxIds 	[]int32 `form:"boxIds"`
	End		bool  `form:"end" ` //游戏是否结束
}

type updateItem struct {
	Type 	int32 `json:"type"`
	Id		int32 `json:"id"`
	Num		int32 `json:"num"`
}

func UpdateGame(c *gin.Context) {

	req := &UpdateGameReq{}
	helper.CheckReq(c, req)
	items := []*updateItem{}
	if err := json.Unmarshal([]byte(req.Items), &items); err != nil {
		log.Debug(err.Error())
	}

	uid := helper.GetClaimUser(c).UID
	for _, item := range items {
		service.AddItem(uid, item.Type, item.Id, item.Num)
	}
	if !req.Guide {
		service.UpdateGameData(uid, req.Score, req.Floor, /*req.Props, req.BoxIds,*/ req.End)
	}
	service.UpdateScore(uid, req.Floor, req.Score, req.Guide)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}

type UpdateBoxReq struct {
	BoxIds []int32 `form:"boxIds"`
}

func UpdateBoxData(c *gin.Context) {
	req := &UpdateBoxReq{}
	helper.CheckReq(c, req)
	service.UpdateBoxIds(helper.GetClaimUser(c).UID, req.BoxIds)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}

type GetGameResp struct {
	InGame	bool	`json:"inGame"`
	Score   int32	`json:"score"`
	Floor	int32	`json:"floor"`
	BoxIds	[]int32	`json:"boxIds"`
}

func GetGame(c *gin.Context) {
	gameData := service.GetGameData(helper.GetClaimUser(c).UID)
	resp := &GetGameResp{
		InGame:gameData.InGame,
		Floor:gameData.Floor,
		Score:gameData.Score,
		BoxIds:gameData.BoxIDs,
		//Props:gameData.Props,
	}
	helper.ResponseWithData(c, resp)
}