package api

import (
	"aliens/log"
	"gangbu/exception"
	"gangbu/module/game/http/helper"
	"gangbu/module/game/service"
	"github.com/gin-gonic/gin"
)

type GetRoleInfoResp struct {
	Score	int32 `json:"score"` //历史最高分
	Floor	int32 `json:"floor"` //历史到达最高关卡
	InGame   bool 	`json:"inGame"` //游戏是否在进行
	Energy	int32	`json:"energy"`
	EnergyLimit int32 `json:"energyLimit"`
	EnergyTime int64 `json:"energyTime"`
	Guide    bool	`json:"guide"`
	LastWatchAd int64 `json:"lastWatchAd"`
	AdTimes 	int32 `json:"adTimes"`
}

func GetRoleInfo(c *gin.Context) {
	role := service.GetRoleInfo(helper.GetClaimUser(c).UID)
	role = service.UpdateEnergy(helper.GetClaimUser(c).UID, false)
	//time.Sleep(3 * time.Second)
	resp := &GetRoleInfoResp{
		Score:	role.Score,
		Floor: 	role.Floor,
		Energy: role.Energy,
		EnergyLimit: role.EnergyLimit,
		EnergyTime:role.EnergyTime.Unix(),
		Guide:role.Guide,
		LastWatchAd:role.LastWatchAd.Unix(),
		AdTimes:role.AdTimes,
	}
	helper.ResponseWithData(c, resp)
}

type GetRankingsReq struct {
	Skip	int32 `form:"skip"`
	Limit	int32 `form:"limit"`

}

type GetRankingsResp struct {
	RankList	[]*service.RankRole		`json:"rankList"`
	Count		int32 `json:"count"`
}

func GetRankings(c *gin.Context) {
	//role := service.GetRoleInfo(helper.GetClaimUser(c).UID)
	req := &GetRankingsReq{}
	helper.CheckReq(c,req)
	rankRoles := service.GetRankings(req.Skip, req.Limit)
	count := service.GetRankCount()
	resp := &GetRankingsResp{
		RankList:rankRoles,
		Count:count,
	}
	helper.ResponseWithData(c, resp)
}

type GetRoleRankReq struct {
	Limit	int32	`form:"limit"`
}

type GetRoleRankResp struct {
	Rank 	int32	`json:"rank"`
	Score 	int32	`json:"score"`
	RankList	[]*service.RankRole		`json:"rankList"` //排名在附近的用户
}

func GetRoleRank(c *gin.Context) {
	req := &GetRoleRankReq{}
	if err := c.ShouldBind(req); err != nil {
		log.Error(err.Error())
		exception.GameException(exception.InvalidParam)
	}
	rank, score, rankList := service.GetRoleRank(helper.GetClaimUser(c).UID, req.Limit)
	resp := &GetRoleRankResp{
		Rank:rank,
		Score:score,
		RankList:rankList,
	}
	helper.ResponseWithData(c, resp)
}

type ReqGetScoreRankReq struct {
	Score	int32 `form:"score"`
}

type ReqGetScoreRankResp struct {
	RankList	[]*service.RankRole		`json:"rankList"` //排名在附近的用户
}

func GetScoreRank(c *gin.Context) {
	req := &ReqGetScoreRankReq{}
	helper.CheckReq(c, req)
	rankList, _ := service.GetLimitRankList(1, req.Score)
	resp := &ReqGetScoreRankResp{
		//Rank:int32(rank),
		RankList:rankList,
	}
	helper.ResponseWithData(c, resp)
}
//type GetRankCountResp struct {
//	Count		int32 `json:"count"`
//}
//
//func GetRankCount(c *gin.Context) {
//	resp := &GetRankCountResp{}
//	resp.Count = service.GetRankCount()
//	helper.ResponseWithData(c, resp)
//}
type UpdateEnergyReq struct {
	IsAd	bool `form:"isAd"`
}

type UpdateEnergyResp struct {
	Energy	int32	`json:"energy"`
	AdTimes int32 `json:"adTimes"`
	EnergyTime int64 `json:"energyTime"`
	LastWatchAd int64 `json:"lastWatchAd"`
}

func UpdateEnergy(c *gin.Context) {
	req := &UpdateEnergyReq{}
	helper.CheckReq(c, req)
	role := service.UpdateEnergy(helper.GetClaimUser(c).UID, req.IsAd)
	resp := &UpdateEnergyResp{
		Energy:	role.Energy,
		EnergyTime: role.EnergyTime.Unix(),
		AdTimes:role.AdTimes,
		LastWatchAd:role.LastWatchAd.Unix(),
	}
	helper.ResponseWithData(c, resp)
}

//type UseEnergyResp struct {
//	Energy	int32	`json:"energy"`
//	EnergyTime int64 `json:"energyTime"`
//}
//
//func UseEnergy(c *gin.Context) {
//	role := service.UseEnergy(helper.GetClaimUser(c).UID)
//	resp := &UseEnergyResp{
//		Energy:role.Energy,
//		EnergyTime:role.EnergyTime.Unix(),
//	}
//	helper.ResponseWithData(c, resp)
//}

func PassGuide(c *gin.Context) {
	service.PassGuide(helper.GetClaimUser(c).UID)
	helper.ResponseWithCode(c, exception.CodeSuccess)
}

//type UpdateScoreReq struct {
//	Floor	int32 `json:"floor"`
//	Score	int32 `json:"score"`
//	End		bool  `json:"end"`
//}

//func UpdateScore(c *gin.Context) {
//	req := &UpdateScoreReq{}
//	if err := c.ShouldBind(req); err != nil {
//		log.Error(err.Error())
//		exception.GameException(exception.InvalidParam)
//	}
//	//service.UpdateScore(helper.GetClaimUser(c).UID, req.Floor, req.Score, req.End)
//	helper.ResponseWithCode(c, exception.CodeSuccess)
//}