package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handlerLogin)
}

func handlerLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	//a := args[1].(gate.Agent)

	log.Debug("%v login", m.UserName)
}
