package internal

import (
	"fmt"
	"server/base"

	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	log.Debug("111111111111111111111111")

	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	s := c.Ref()
	err = s.DB("test").C("counters").RemoveId("test")
}

func (m *Module) OnDestroy() {

}
