package db

import (
	"aliens/common/util"
	"gangbu/constant"
	"gangbu/exception"
	"gangbu/module/game/service/lpc"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func UpdateOne(data interface{}) {
	err := databaseHandler.UpdateOne(data)
	if err != nil {
		exception.GameExceptionCustom("Update", exception.DatabaseError, err)
	}
}

func InsertOne(api string, data interface{}) {
	err := databaseHandler.Insert(data)
	if err != nil {
		exception.GameExceptionCustom(api, exception.DatabaseError, err)
	}
}

func GetUserInfo(username string) *DBUser {
	user := &DBUser{}
	if err := databaseHandler.QueryOneCondition(user, "username", username); err != nil {
		if err.Error() == "not found" {
			return nil
		} else {
			exception.GameExceptionCustom("GetUserInfo", exception.DatabaseError, err)
		}
	}
	return user
}

func GetUserByUid(uid int32) *DBUser {
	user := &DBUser{}
	if err := databaseHandler.QueryOneCondition(user, "_id", uid); err != nil {
		if err.Error() == "not found" {
			log.Debug("user not found, uid: %v", uid)
			return nil
		} else {
			exception.GameExceptionCustom("GetRoleByUid", exception.DatabaseError, err)
		}
	}
	return user
}

func NewUser(username, password, openID, avatar, nickname string, channel int32) *DBUser {
	user := &DBUser{
		Username: username,
		Nickname:nickname,
		Channel:channel,
		Phone:    "",
		OpenID:   openID,
		Status:   0,
		Avatar:   avatar,
		RegTime:  time.Now(),
	}
	InsertOne("NewUser", user)
	return user
}

func DeleteUser(uid int32) {
	err := databaseHandler.DeleteOneCondition(&DBUser{}, bson.D{bson.DocElem{Name:"_id",Value:uid}})
	if err != nil {
		exception.GameExceptionCustom("DeleteUser", exception.DatabaseError, err)
	}
}

//func UpdateUserWechatInfo(uid int32, nickname, avatar string) error {
//	udoc := bson.M{"nickname":nickname, "avatar":avatar}
//	return updateUserByUid("user", uid, udoc)
//}



//func updateUserByUid(collection string, uid int32, udoc bson.M) error {
//	qdoc := bson.M{"_id":uid}
//	return databaseHandler.Update(collection, qdoc, bson.M{"$set":udoc})
//}

//更新活跃时间 用来查询新增留存和活跃留存
func (this *DBUser)UpdateActiveTime() {
	this.LastActiveTIme = this.ActiveTime
	this.ActiveTime = time.Now()
	//regTime := this.RegTime.Local().Format("2006-01-02")
}

func (this *DBUser)LoginLog (new bool, channel int32) {
	day := time.Now().Local().Format("2006-01-02")
	exist, err := databaseHandler.IDExist(&DBLoginLog{ID:&CustomLoginLogID{UID:this.ID, Day:day}})
	if err != nil {
		exception.GameExceptionCustom("LoginLog", exception.DatabaseError, err)
	}
	if !exist {
		lpc.DBServiceProxy.Insert(&DBLoginLog{
			ID:      &CustomLoginLogID{UID:this.ID, Day:day},
			New:     new,
			Channel: channel,
		},databaseHandler)
	}
}

//-----------------------------------------role--------------------------------------------------
func GetRoleByUid(uid int32) *DBRole {
	role := &DBRole{}
	if err := databaseHandler.QueryOneCondition(role, "uid", uid); err != nil {
		//log.Debug("err:%v",err.Error())
		if err.Error() == "not found" {
			return nil
		} else {
			exception.GameExceptionCustom("GetRoleByUid", exception.DatabaseError, err)
		}
	}
	return role
}

func NewRoleByUid(uid int32) *DBRole {
	role := &DBRole{
		UID:uid,
		Energy:constant.ENERGY_INIT,
		EnergyLimit:constant.ENERGY_LIMIT,
		EnergyTime:time.Now(),
		AdTimes:constant.MAX_DAY_AD_RESTORE,
	}
	InsertOne("NewRoleByUid", role)
	return role
}

//func (this *DBRole) Update() error {
//	return databaseHandler.UpdateOne(this)
//}
func DeleteRole(uid int32) {
	err := databaseHandler.DeleteOneCondition(&DBRole{}, bson.D{bson.DocElem{Name:"uid",Value:uid}})
	if err != nil {
		exception.GameExceptionCustom("DeleteRole", exception.DatabaseError, err)
	}
}
//--------------------------------------------energy-------------------------------------------------
func (this *DBRole) TakeInEnergy(energy int32, limit bool) bool {
	if energy <= 0 {
		return false
	}
	resultEnergy := this.Energy + energy

	energyLimit := this.EnergyLimit
	if limit && resultEnergy > energyLimit {
		if this.Energy < energyLimit {
			this.Energy = energyLimit
		} else {
			return false
		}
	} else {
		this.Energy = resultEnergy
	}
	return true
}

func (this *DBRole) TakeOutEnergy(energy int32) bool {
	oldEnergy := this.Energy

	if oldEnergy < energy {
		return false
	}
	this.Energy -= energy
	if this.Energy >= this.EnergyLimit {
		this.EnergyTime = time.Now()
	}
	return true
}

func (this *DBRole) UpdateHelpTime() {
	timestamp := util.GetTodayHourTime(0)
	if this.LastHelpTime.Before(timestamp) {
		this.TodayHelper = []int32{}
	}
}

func (this *DBRole) CheckHelp(helpID int32) bool {
	if len(this.TodayHelper) >= constant.MAX_DAY_SHARE_HELP_NUM {
		return false
	}
	if util.ContainsInt32(helpID, this.TodayHelper) {
		return false
	}
	this.LastHelpTime = time.Now()
	this.TodayHelper = append(this.TodayHelper, helpID)
	return true
}

//----------------------------------------rank--------------------------------------------
//分页查询排行榜
func GetRolesRank(skip int32, limit int32) []*DBRole {
	//query := make(bson.M)
	var roles []*DBRole
	if err := databaseHandler.QueryAllConditionsSkipLimit(&DBRole{}, bson.M{"score":bson.M{"$gt":0}}, &roles, int(skip * limit), int(limit), "-score"); err != nil {
		exception.GameExceptionCustom("GetRolesRank", exception.DatabaseError, err)
	}
	return roles
}
//排行榜个数
func GetRankCount () int32 {
	count, err := databaseHandler.QueryConditionsCount(&DBRole{}, bson.M{"score":bson.M{"$gt":0}})
	if  err != nil {
		exception.GameExceptionCustom("GetRankCount", exception.DatabaseError, err)
	}
	return int32(count)
}

//查询玩家当前排名
func GetScoreRank(score int32) int {
	query := make(bson.M)
	query["score"] = bson.M{"$gt":score}
	rank, err := databaseHandler.QueryConditionsCount(&DBRole{}, query)
	if err != nil {
		exception.GameExceptionCustom("GetScoreRank", exception.DatabaseError, err)
	}
	return rank
}

func GetScoreLimitRank(score int32, limit int32) ([]*DBRole, []*DBRole) {
	var gt_roles []*DBRole
	var lt_roles []*DBRole
	query := make(bson.M)
	query["score"] = bson.M{"$gt":score}
	err1 := databaseHandler.QueryAllConditionsLimit(&DBRole{}, query, &gt_roles, int(limit), "+score")
	if err1 != nil {
		exception.GameExceptionCustom("GetScoreLimitRank", exception.DatabaseError, err1)
	}
	query["score"] = bson.M{"$lt":score}
	err2 := databaseHandler.QueryAllConditionsLimit(&DBRole{}, query, &lt_roles, int(limit), "-score")
	if err2 != nil {
		exception.GameExceptionCustom("GetScoreLimitRank", exception.DatabaseError, err2)
	}
	return gt_roles, lt_roles
}

func (this *DBRole) UpdateRoleScore(score int32, floor int32) {
	if this.Score < score {
		this.Score = score
	}
	if this.Floor < floor {
		this.Floor = floor
	}
}

//func (this *DBRole)UpdateFloorScore(score int32, floor int32) {
//	if this.Score < score {
//		this.Score = score
//	}
//	if this.Floor < floor {
//		this.Floor = floor
//	}
//}

//-------------------------------------prop--------------------------------------
//
//func (this *DBProp) TakeIn(num int32) {
//	this.Num += num
//}
//
//func (this *DBProp) TakeOut(num int32) bool {
//	if this.Num < num {
//		return false
//	}
//	this.Num -= num
//	return true
//}
//
//func (this *DBProp) Update() {
//	err := databaseHandler.UpdateOne(this)
//	if err != nil {
//		exception.GameExceptionCustom("PropUpdate", exception.DatabaseError, err)
//	}
//}
//
//func GetPropsByUid(uid int32) []*DBProp {
//	var props []*DBProp
//	if err := databaseHandler.QueryAllCondition(&DBProp{}, "_id.uid", uid, &props); err != nil {
//		exception.GameExceptionCustom("GetPropsByUid", exception.DatabaseError, err)
//	}
//	return props
//}
//
//func GetPropByUidAndType(uid int32, PropType int32) *DBProp {
//	prop := &DBProp{}
//	if err := databaseHandler.QueryOneConditions(prop, bson.M{"_id.uid":uid, "_id.type":PropType}); err != nil {
//		if err.Error() == "not found" {
//			return nil
//		} else {
//			exception.GameExceptionCustom("GetPropByUidAndType", exception.DatabaseError, err)
//		}
//	}
//	return prop
//}
//
//func CreateProp(uid int32, PropType int32, propNum int32){
//	prop := &DBProp{
//		ID:&CustomID{
//			UID:uid,
//			Type:PropType,
//		},
//		Num:propNum,
//	}
//	if err := databaseHandler.Insert(prop); err != nil {
//		exception.GameExceptionCustom("CreateProp", exception.DatabaseError, err)
//	}
//}

func AddInviterPropByUid(uid int32, PropType int32) {
	//databaseHandler
}

//-------------------------------------------game----------------------------------------------

//func (this *DBGameData)Init(props []*Prop) {
//	this.InGame = true
//	//this.Props =
//	this.Score = 0
//	this.Floor = 0
//	if props == nil || len(props) == 0 {
//		return
//	}
//	for index, prop := range props {
//		this.Props[index] = &Prop{Type:prop.Type, Num:prop.Num}
//	}
//}

func (this *DBGameData) Clean() {
	this.Floor = 0
	this.Score = 0
	//this.Props = []*Prop{}
	this.BoxIDs = []int32{}
	this.InGame = false
}

func (this *DBGameData) Update(floor int32, score int32/*, props []*Prop, boxIDs []int32*/) {
	this.InGame = true
	this.Floor = floor
	this.Score = score
	//for _, prop := range props {
	//	this.Props = append(this.Props, prop)
	//}
	//this.BoxIDs = boxIDs
}

func NewGameData(uid int32) *DBGameData {
	gameData := &DBGameData{
		UID:uid,
		Floor:0,
		Score:0,
		//Props:[]*Prop{},
		BoxIDs:[]int32{},
		InGame:false,
	}
	InsertOne("NewGameData", gameData)
	return gameData
}

func GetGameDataByUid(uid int32) *DBGameData {
	gameData := &DBGameData{}
	if err := databaseHandler.QueryOneCondition(gameData, "uid", uid); err != nil {
		log.Debug("err:%v",err.Error())
		if err.Error() == "not found" {
			return nil
		} else {
			exception.GameExceptionCustom("GetGameDataByUid", exception.DatabaseError, err)
		}
	}
	return gameData
}


//-------------------------------------------notice------------------------------------------------

func GetCurrentNotice() *DBNotice {
	notice := []*DBNotice{}
	query := bson.M{"time":bson.M{"$lt":time.Now()}}
	err := databaseHandler.QueryAllConditionsLimit(&DBNotice{}, query, &notice, 1, "-time")
	if err!=nil {
		exception.GameExceptionCustom("GetCurrentNotice", exception.DatabaseError, err)
	}
	if len(notice) != 1 {
		return nil
	}
	return notice[0]
}

func CreateNotice(title, context string, pubTime time.Time) {
	notice := &DBNotice{
		Title:   title,
		Content: context,
		PubTime: pubTime,
	}
	InsertOne("CreateNotice", notice)
}

/*-----------------------------------------item------------------------------------------*/
func GetItemByID(uid, itemType, id int32) *DBItem {
	item := &DBItem{}
	if err := databaseHandler.QueryOneConditions(item, bson.M{"_id.uid":uid, "_id.type": itemType, "_id.id":id}); err != nil {
		if err.Error() == "not found" {
			return nil
		} else {
			exception.GameExceptionCustom("GetPropByUidAndType", exception.DatabaseError, err)
		}
	}
	return item
}

func CreateItem(uid, itemType, itemId, num int32) *DBItem {
	item := &DBItem{
		ID:&CustomID{UID:uid, Type:itemType, ID:itemId},
		Num:num,
	}
	InsertOne("CreateItem", item)
	return item
}

func GetItemsByType(uid, itemType int32) []*DBItem {
	items := []*DBItem{}
	if err := databaseHandler.QueryAllConditions(&DBItem{},bson.M{"_id.uid":uid, "_id.type": itemType}, &items); err != nil {
		exception.GameExceptionCustom("GetPropsByUid", exception.DatabaseError, err)
	}
	return items
}

func (this *DBItem) TakeIn(num int32) {
	this.Num += num
}

func (this *DBItem) TakeOut(num int32) bool {
	if this.Num < num {
		return false
	}
	this.Num -= num
	return true
}

//--------------------------statistic---------------------------//
func QueryCount (query bson.M) int32 {
	//query := make(bson.M)
	//query[field] = bson.M{"$gte":start, "$lte":end}
	count, err := databaseHandler.QueryConditionsCount(&DBUser{}, query)
	if err != nil {
		exception.GameException(exception.DatabaseError)
	}
	return int32(count)
}

func QuerysCount (field string, start time.Time, end time.Time) int32 {
	query := make(bson.M)
	query[field] = bson.M{"$gte":start, "$lte":end}
	count, err := databaseHandler.QueryConditionsCount(&DBUser{}, query)
	if err != nil {
		exception.GameException(exception.DatabaseError)
	}
	return int32(count)
}

func AppendQuery (query bson.M, field string,  start time.Time, end time.Time) {
	if query == nil {
		query = make(bson.M)
	}
	query[field] = bson.M{"$gte":start, "$lte":end}
}