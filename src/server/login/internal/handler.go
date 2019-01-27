package internal

import (
	"reflect"
	mongodbmgr "server/db"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var (
	Agents map[int]*gate.Agent
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.CSLogin{}, handlerLogin)
	Agents = make(map[int]*gate.Agent)
}

func handlerLogin(args []interface{}) {
	m := args[0].(*msg.CSLogin)
	a := args[1].(gate.Agent)

	log.Debug("%v login", m.UserName)
	//mychan.Dosomething()

	ud, state := mongodbmgr.Login(m.UserName)
	if state == 0 {
		a.SetUserData(ud)
		a.WriteMsg(&msg.SCLogin{
			ErrorCode: 0,
			UserId:    ud.UserId,
		})
	} else {
		a.SetUserData(ud)
		a.WriteMsg(&msg.SCLogin{
			ErrorCode: 0,
			UserId:    ud.UserId,
		})
	}

	Agents[ud.UserId] = &a
	log.Debug("login success and insert success %v", ud)

}
