package constant

const (
	ENERGY_INIT			= 6
	ENERGY_LIMIT		= 6
	ENERGY_COST			= 1
	ENERGY_RESTORE_TIME	= 10*60 //s
	ENERGY_RESTORE_NUM	= 1
	ENERGY_AD_RESTORE	= 1

	WATCH_AD_INTERVAL  = 60*60 //s
	MAX_DAY_AD_RESTORE = 6     //每天看广告恢复体力的次数

	FLOOR_FIRST			= 1
	MAX_FLOOR_SCORE		= 100000

	PROP_PER_COST = 1 //每次道具消耗的数量

	RANK_MAX_LIMIT = 5

	MAX_DAY_SHARE_HELP_NUM = 5
)

const (
	PROP_TYPE_HP		= 	1 //血
	PROP_TYPE_SHIELD	=	2 //盾
	PROP_TYPE_RELIVE	=	3 //复活
	PROP_TYPE_CLEAN		=	4 //清除debuff
)

const (
	PLATFORM_WECHAT = "wx"
	PLATFORM_TOUTIAO = "tt"
	PLATFORM_VISITOR = "visitor"
)

const (
	CHANNEL_WECHAT  = 1
	CHANNEL_TOUTIAO = 2
	CHANNEL_VISITOR = 3
	CHANNEL_WEB     = 4
)

const (
	ITEM_PROP 		= 1	//道具(药水) (both)
	ITEM_RELIC 		= 2 //碎片 (冈布奥)
	ITEM_EQUIPMENT 	= 3 //装备 (冈布奥)
	ITEM_ARCHIVE 	= 4 //伴手礼 (赛几)
)

const (
	DB_COMMAND_INSERT = "I"
	DB_COMMAND_UPDATE = "U"
	DB_COMMAND_DELETE = "D"
	DB_COMMAND_FUPDATE = "FU"
	DB_COMMAND_CONDITION_UPDATE = "CU"
	DB_COMMAND_CONDITION_DELETE = "CD"

	DB_DEBUG                 = true	 //是否开启数据库操作
	DB_SIGAL_TIMEOUT float64 = 1 //数据库操作超时告警阈值 1秒
)