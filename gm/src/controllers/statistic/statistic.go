package statistic

import (
	"db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func GetStatisticNewlyCount(field string, appendField string, fieldTime string, appendFieldTime string) (fieldCount int32, appendFieldCount int32){
	//reg_start, reg_end := ParseTime(fieldTime)
	//active_start, active_end := ParseTime(appendFieldTime)
	//query := make([]bson.M,0)
	//AppendQuery(query, field, reg_start, reg_end)
	query := make(bson.M)
	query["_id.day"] = fieldTime
	query["new"] = true
	fieldCount = db.QueryCount(query)
	//AppendQuery(query, appendField, active_start, active_end)
	pipe := []bson.M{
		{"$match":bson.M{"_id.day":fieldTime, "new":true}},
		{"$lookup":bson.M{"from":"loginLog","localField":"_id.uid","foreignField":"_id.uid","as":"leftJoin"}},
		{"$match":bson.M{"leftJoin":bson.M{"$elemMatch":bson.M{"_id.day":appendFieldTime,"new":false}}}},
	}
	appendFieldCount = db.PipeCount(pipe)
	return fieldCount, appendFieldCount
}

func GetStatisticActivityCount(field string, appendField string, fieldTime string, appendFieldTime string) (fieldCount int32, appendFieldCount int32){

	query := make(bson.M)
	query["_id.day"] = fieldTime
	fieldCount = db.QueryCount(query)
	pipe := []bson.M{
		{"$match":bson.M{"_id.day":fieldTime}},
		{"$lookup":bson.M{"from":"loginLog","localField":"_id.uid","foreignField":"_id.uid","as":"leftJoin"}},
		{"$match":bson.M{"leftJoin":bson.M{"$elemMatch":bson.M{"_id.day":appendFieldTime}}}},
	}
	appendFieldCount = db.PipeCount(pipe)
	return fieldCount, appendFieldCount
}

func ParseTime(reg string) (time.Time, time.Time) {
	t, err :=  time.ParseInLocation("2006-01-02", reg, time.Local)
	if err != nil {
		//exception.GameException(exception.TimeParseError)
	}
	year, month, day := t.Date()
	t_start := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	t_end := time.Date(year, month, day, 23, 59 ,59, 59, t.Location())
	return t_start, t_end
}

func AppendQuery (query bson.M, field string,  start time.Time, end time.Time) {
	if query == nil {
		query = make(bson.M)
	}
	query[field] = bson.M{"$gte":start, "$lte":end}
}

func GetUserStatistic (day string, new bool) int32 {
	return db.QueryCount(bson.M{"_id.day":day, "new":new})
}

func GetUserCount() int32 {
	return db.QueryCount(bson.M{})
}