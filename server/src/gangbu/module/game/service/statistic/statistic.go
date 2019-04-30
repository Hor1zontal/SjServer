package statistic

import (
	"gangbu/exception"
	"gangbu/module/game/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//func GetNewlyCount(reg string, active string) (int32, int32){
//	reg_start, reg_end := ParseTime(reg)
//	active_start, active_end := ParseTime(active)
//	var query bson.M
//	db.AppendQuery(query, "reg", reg_start, reg_end)
//	regCount := db.QueryCount(query)
//	db.AppendQuery(query, "active", active_start, active_end)
//	activeCount := db.QueryCount(query)
//	return regCount, activeCount
//}

func GetStatisticCount(field string, appendField string, fieldTime string, appendFieldTime string) (fieldCount int32, appendFieldCount int32){
	reg_start, reg_end := ParseTime(fieldTime)
	active_start, active_end := ParseTime(appendFieldTime)
	query := make(bson.M)
	db.AppendQuery(query, field, reg_start, reg_end)
	fieldCount = db.QueryCount(query)
	db.AppendQuery(query, appendField, active_start, active_end)
	appendFieldCount = db.QueryCount(query)
	return fieldCount, appendFieldCount
}


func ParseTime(reg string) (time.Time, time.Time) {
	t, err :=  time.ParseInLocation("2006-01-02", reg, time.Local)
	if err != nil {
		exception.GameException(exception.TimeParseError)
	}
	year, month, day := t.Date()
	t_start := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	t_end := time.Date(year, month, day, 23, 59 ,59, 59, t.Location())
	return t_start, t_end
}