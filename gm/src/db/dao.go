package db

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func QueryCount(query bson.M) int32 {
	count, err := databaseHandler.QueryConditionsCount(&DBLoginLog{}, query)
	//result := []*DBLoginLog{}
	//databaseHandler.QueryAllConditions(&DBLoginLog{}, bson.M{}, &result)
	if err != nil {
		fmt.Println("query count error", err.Error())
	}
	return int32(count)
}

func PipeCount(pipeline []bson.M) int32 {
	result := []bson.M{}
	err := databaseHandler.PipeAllConditions(&DBLoginLog{}, pipeline, &result)
	//count, err := databaseHandler.QueryConditionsCount(&DBLoginLog{}, query)
	if err != nil {
		fmt.Println("query count error", err.Error())
	}
	return int32(len(result))
}