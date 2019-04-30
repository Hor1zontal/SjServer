package service

import (
	"gangbu/exception"
	"gangbu/module/game/db"
)

func AddHelp(helpID, beHelpId int32) bool {
	role := db.GetRoleByUid(beHelpId)
	if role == nil {
		exception.GameException(exception.RoleNotFound)
	}
	role.UpdateHelpTime()
	if !role.CheckHelp(helpID) {
		return false
	}
	db.UpdateOne(role)
	return true
}

//-----------------------------------------------------------------------------------------------------------------
type Item struct {
	ID		int32 `json:"id" form:"id"`
	Num 	int32 `json:"num" form:"num"`
}

func AddItem(uid, itemType, itemID, num int32) {
	item := db.GetItemByID(uid, itemType, itemID)
	if item == nil {
		item = db.CreateItem(uid, itemType, itemID, num)
	}
	item.TakeIn(num)
	db.UpdateOne(item)
}

func UseItem(uid, itemType, itemId, itemNum int32) {
	item := db.GetItemByID(uid, itemType, itemId)
	if item == nil {
		exception.GameException(exception.PropNotEnough)
	}
	if !item.TakeOut(itemNum) {
		exception.GameException(exception.PropNotEnough)
	}
	db.UpdateOne(item)
}


func GetItems(uid, itemType int32) []*Item{
	items_db := db.GetItemsByType(uid, itemType)
	result := make([]*Item,len(items_db))
	for index, item_db := range items_db {
		result[index] = &Item{ID:item_db.ID.ID,Num:item_db.Num}
	}
	return result
}