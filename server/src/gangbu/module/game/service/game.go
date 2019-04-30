package service

import (
	"gangbu/constant"
	"gangbu/exception"
	"gangbu/module/game/db"
	"gangbu/module/game/service/lpc"
	"time"
)

func EnterGameData(uid int32/*, props []*db.Prop*/) {
	//初始化，数据清空
	gameData := db.GetGameDataByUid(uid)
	if gameData == nil {
		exception.GameException(exception.GameDataNotFound)
	}
	//初始化带入游戏的道具
	if !gameData.InGame {
		////扣掉道具
		//for _, v := range props {
		//	prop := db.GetPropByUidAndType(uid, v.Type)
		//	if prop == nil {
		//		exception.GameException(exception.PropNotFound)
		//	}
		//	if !prop.TakeOut(prop.Num) {
		//		exception.GameException(exception.PropNotEnough)
		//	}
		//	prop.Update()
		//}
		//初始化带入游戏的道具
		//gameData.Init(props)
		gameData.Clean()
		gameData.InGame = true
		db.UpdateOne(gameData)
	}
}

func UpdateGameData(uid int32, score int32, floor int32, /*props []*db.Prop, boxIDs []int32,*/ end bool) {
	if score < 0 {
		exception.GameException(exception.UpdateScoreError)
	}
	gameData := db.GetGameDataByUid(uid)
	if gameData == nil {
		exception.GameException(exception.GameDataNotFound)
	}
	if score - gameData.Score < 0 || score - gameData.Score > constant.MAX_FLOOR_SCORE {
		exception.GameException(exception.UpdateScoreError)
	}
	if floor != gameData.Floor + 1 {
		//不是下一层
		exception.GameException(exception.UpdateScoreError)
	}
	//log.Debug("lpc insert in DBGameLog")
	lpc.DBServiceProxy.Insert(&db.DBGameLog{
		UID:uid,
		End:end,
		Floor:floor,
		BoxIDs:gameData.BoxIDs,
		Score:score,
		Time:time.Now(),
	}, db.Database.GetHandler())

	if end {
		//游戏结束 游戏数据清零
		gameData.Clean()
	} else {
		//更新分数、道具等信息
		gameData.Update(floor, score/*, boxIDs, props, */)
	}
	db.UpdateOne(gameData)
}

func GetGameData(uid int32) *db.DBGameData {
	gameData := db.GetGameDataByUid(uid)
	if gameData == nil {
		gameData = db.NewGameData(uid)
	}
	return gameData
}

func UpdateBoxIds(uid int32, boxIds []int32) {
	gameData := db.GetGameDataByUid(uid)
	if gameData == nil {
		exception.GameException(exception.GameDataNotFound)
	}
	//gameData.BoxIDs = []int32{}
	//for _, id := range boxIds {
	//	if id != 0 {
	//		gameData.BoxIDs = append(gameData.BoxIDs, id)
	//	}
	//}
	gameData.BoxIDs = boxIds
	db.UpdateOne(gameData)
}