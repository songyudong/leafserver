package internal

import (
	"reflect"
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

	ud := &msg.UserData{
		UserId:   35678,
		UserName: m.UserName,
	}
	a.SetUserData(ud)
	a.WriteMsg(&msg.SCLogin{
		ErrorCode: 0,
		UserId:    35678,
	})

	Agents[35678] = &a
}
