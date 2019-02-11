package internal

import (
	"reflect"
	
	
	"test/message"
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
	handleMsg(&message.Hello{}, handlerHello)
	Agents = make(map[int]*gate.Agent)
}

func handlerHello(args []interface{}) {
	m := args[0].(*message.Hello)
	_ = args[1].(gate.Agent)

	log.Debug("%v login", m.Name)

}
