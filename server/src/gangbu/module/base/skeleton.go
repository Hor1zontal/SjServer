package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
)

const (
	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              GoLen,
		TimerDispatcherLen: TimerDispatcherLen,
		AsynCallLen:        AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}

