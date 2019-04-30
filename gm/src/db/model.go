package db

import "time"

type DBLoginLog struct {
	ID	*CustomLoginLogID `bson:"_id"`
	New bool `bson:"new"`
	Channel int32 `bson:"channel"`
}

type CustomLoginLogID struct {
	UID int32 	`bson:"uid"`
	Day string 	`bson:"day"` //活跃时间（某天）
}

type DBNotice struct {
	ID      int32     `bson:"_id" gorm:"AUTO_INCREMENT"`
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
	PubTime time.Time `bson:"time"`
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
