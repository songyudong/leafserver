package internal

import (
	"reflect"
	"server/login"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Hello{}, handlerHello)
	handler(&msg.CSChat{}, handlerChat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.Name)
	a.WriteMsg(&msg.Hello{
		Name: "client",
	})
}

func handlerChat(args []interface{}) {
	m := args[0].(*msg.CSChat)
	log.Debug("c2s chat %v", m.Content)

	for k, v := range login.Agents {
		log.Debug("send to client: %v %v %v", k, v, m.Content)
		a := *v
		ud := a.UserData().(*msg.UserData)
		log.Debug("message : id = %v name = %v", ud.UserId, ud.UserName)
		a.WriteMsg(&msg.SCChat{
			UserId:   ud.UserId,
			UserName: ud.UserName,
			Content:  m.Content,
		})
	}
}
