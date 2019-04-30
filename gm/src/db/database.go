package db

import (
	"aliens/database"
	"aliens/database/dbconfig"
	"aliens/database/mongo"
	"github.com/astaxie/beego"
)

// 连接mongodb数据库
var (
	MongodbAddr   string = "" //mongodb数据库地址
	MongodbName   string = "" //mongodb数据名称
	//MongodbUser   string = "" //mongodb用户名
	//MongodbPasswd string = "" //mongodb密码
)

var Database database.IDatabase = &mongo.Database{}
var databaseHandler = Database.GetHandler()

func init() {
	MongodbAddr = beego.AppConfig.String("mongodb_addr")
	MongodbName = beego.AppConfig.String("mongodb_name")
	//MongodbUser = beego.AppConfig.String("mongodb_username")
	//MongodbPasswd = beego.AppConfig.String("mongodb_passwd")
}

func Init() {

	err := Database.Init(dbconfig.DBConfig{
		Name:MongodbName,
		Address:MongodbAddr,
	})
	if err != nil {
		panic(err)
	}

	databaseHandler.EnsureTable("loginLog", &DBLoginLog{})
	databaseHandler.EnsureTable("notice", &DBNotice{})
	databaseHandler.EnsureTable("gameLog", &DBGameLog{})
}