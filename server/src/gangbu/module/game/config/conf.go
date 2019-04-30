package config

import (
	"aliens/database/dbconfig"
	"aliens/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	//Debug = false
	//Tag   = ""
	Root  = ""
)

var Server *BaseConfig

type BaseConfig struct {

	Platform	string 				`yaml:"platform"`
	//Enable       bool				`yaml:"enable"`
	Database     dbconfig.DBConfig  `yaml:"database"`

	HTTPAddress  string				`yaml:"http"`
	WxConfig	Config				`yaml:"wxConfig"`
	TtConfig	Config				`yaml:"ttConfig"`

	JWTSecret	 string				`yaml:"JWTSecret"`
	JWTExpired	 int64				`yaml:"JWTExpired"`
	IsSign		 bool				`yaml:"isSign"`
	ExpiredTime  int32				`yaml:"expiredTime"`

	IsCors		bool 				`yaml:"isCors"`
	//DefaultPWD string
}

type Config struct {
	AppID string			`yaml:"AppID"`
	AppSecret string		`yaml:"AppSecret"`
}

func Init(rootPath string) {
	Root = rootPath
	LoadConfigData("/config.yml", &Server)
}

func LoadConfigData(path string, config interface{}) {
	LoadConfigDataEx(path, config,true)
}

func GetConfigPath(path string) string {
	return Root + path
}

func LoadConfigDataEx(path string, config interface{}, fatal bool) {
	if config == nil {
		return
	}
	path = Root + path
	data, err := ioutil.ReadFile(path)
	if err != nil {
		if fatal {
			log.Fatal("config file %v  is not found", path)
		}
		return
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		if fatal {
			log.Fatal("load config %v err %v", path, err)
		}
	}
}
