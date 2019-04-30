package db

import (
	"time"
)

type DBUser struct {
	ID		 int32  `bson:"_id" gorm:"AUTO_INCREMENT"`		//用户id
	Username string `bson:"username" unique:"true"`			//用户名 渠道信息_渠道用户id存Username
	Nickname string `bson:"nickname"`
	//Password string `bson:"passport"`						//加密的密码
	//Salt     string `bson:"salt"`							//加密的salt

	OpenID string `bson:"openid"` //微信OPENID 绑定微信填写
	//Token    string `bson:"token"`

	//ChannelUID string `bson:"cuid"`    //用户的渠道的渠道用户id
	//Channel    string `bson:"channel"` //用户的渠道信息 渠道用户id存Username
	Avatar	   string `bson:"avatar"` //用户的头像地址

	Phone string `bson:"phone"` //用户电话
	Platform string `bson:"platform"` //平台
	Channel int32 `bson:"channel"` // 渠道
	//IP     string `bson:"ip"`     //最后一次登录的ip
	Status  byte      `bson:"status"`  //用户状态 0正常  1封号
	RegTime time.Time `bson:"regtime"` //用户注册时间
	ActiveTime time.Time `bson:"activeTime"` //最后一次登录时间
	LastActiveTIme time.Time `bson:"lastActiveTime"` //上一次登录时间

}

////玩家冈布奥的加成
//type DBHero struct {
//	ID		*CustomID	`bson:"_id"`
//	//Type    int32  		`bson:"type"`
//}


type DBRole struct {
	ID		int32 `bson:"_id" gorm:"AUTO_INCREMENT"`
	UID		int32 `bson:"uid" unique:"true"`
	Score	int32 `bson:"score" unique:"false"` //历史最高分
	Floor	int32 `bson:"floor"` //历史到达最高关卡
	//CurrScore int32 `bson:"cScore"` //当前分数
	//CurrFloor int32 `bson:"cFloor"` //当前关卡数
	Energy	int32 `bson:"energy"` //体力
	EnergyLimit int32 `bson:"energyLimit"` //体力上限
	EnergyTime time.Time `bson:"energyTime"` //体力上次的恢复时间戳
	LastWatchAd time.Time `bson:"lastWatchAd"` //上次看广告恢复体力的时间
	AdTimes int32 `bson:"adTimes"` //当天通过看广告恢复体力的次数
	Guide	bool `bson:"guide"` //是否是通过新手引导
	LastHelpTime time.Time `bson:"lastHelpTime"`
	TodayHelper []int32 `bson:"todayHelper"`
}

type DBLoginLog struct {
	//ID	int32 `bson:"_id" gorm:"AUTO_INCREMENT"`
	ID	*CustomLoginLogID `bson:"_id"`
	//UID int32 `bson:"uid" unique:"false"`
	//Register string `bson:"register" unique:"false"` //注册时间（某天）
	//Active string `bson:"active" unique:"false"` //活跃时间（某天）
	New     bool  `bson:"new"`
	Channel int32 `bson:"channel"`
}

type CustomLoginLogID struct {
	UID int32 `bson:"uid"`
	Day string `bson:"day"` //活跃时间（某天）
}

//游戏中的数据
type DBGameData struct {
	ID		int32 `bson:"_id" gorm:"AUTO_INCREMENT"`
	UID		int32 `bson:"uid" unique:"true"`
	InGame	bool `bson:"inGame"` //游戏是否在进行
	Floor	int32 `bson:"floor"`
	Score	int32 `bson:"score"`
	//Props 	[]*Prop `bson:"props"`
	BoxIDs	[]int32 `bson:"boxIDs"`
}

type DBGameLog struct {
	ID		int32 `bson:"_id" gorm:"AUTO_INCREMENT"`
	UID		int32 `bson:"uid" unique:"false"`
	Floor   int32 `bson:"floor"`
	Score	int32 `bson:"score"`
	BoxIDs  []int32 `bson:"boxIds"`
	End  	bool `bson:"end"`
	Time    time.Time `bson:"time"`
}

//玩家的道具
type DBItem struct {
	ID *CustomID `bson:"_id"`
	Num int32 `bson:"num"`
}

type DBCheck struct {
	ID			int32		`bson:"_id" gorm:"AUTO_INCREMENT"`
	UID		 	int32		`bson:"uid" unique:"true"`
	CheckTime	[]time.Time	`bson:"checkTime"`
}

type DBMail struct {
	ID			int64		`bson:"_id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
}

type CustomID struct {
	UID		int32 `bson:"uid"`
	Type	int32 `bson:"type"`
	ID		int32 `bson:"id"`
}

type DBNotice struct {
	ID      int32     `bson:"_id" gorm:"AUTO_INCREMENT"`
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
	PubTime time.Time `bson:"time"`
}