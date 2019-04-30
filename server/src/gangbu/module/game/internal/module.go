package internal

import (
	"gangbu/module/base"
	"gangbu/module/game/db"
	"gangbu/module/game/http"
	"github.com/name5566/leaf/module"
)

var(
	skeleton = base.NewSkeleton()
)

type Module struct {
	*module.Skeleton
}

func (m *Module) IsEnable() bool {
	return true
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	db.Init()
	http.Init()
}

func (m *Module) OnDestroy() {
	http.Close()
	db.Close()
}

