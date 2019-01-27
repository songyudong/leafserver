package internal

import (
	"reflect"
	mongodbmgr "server/db"
	"server/login"
	"server/msg"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

var (
	Clients map[int]*chanrpc.Client
)

func init() {
	handler(&msg.Hello{}, handlerHello)
	handler(&msg.CSChat{}, handlerChat)
	handler(&msg.CSMatch{}, handlerMatch)
	handler(&msg.CSEnterGame{}, handlerEnterGame)

	Clients = make(map[int]*chanrpc.Client)
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
		ud := a.UserData().(*mongodbmgr.DBUser)
		log.Debug("message : id = %v name = %v", ud.UserId, ud.UserName)
		a.WriteMsg(&msg.SCChat{
			UserId:   ud.UserId,
			UserName: ud.UserName,
			Content:  m.Content,
		})
	}
}

func handlerMatch(args []interface{}) {
	m := args[0].(*msg.CSMatch)
	a := args[1].(gate.Agent)
	log.Debug("c2s chat mode %v", m.Mode)

	RoomId := NewRoom()

	a.WriteMsg(&msg.SCMatch{
		Result: 0,
		Room:   RoomId,
	})
}

func handlerEnterGame(args []interface{}) {
	m := args[0].(*msg.CSEnterGame)
	a := args[1].(gate.Agent)
	log.Debug("c2s chat room %v", m.Room)

	a.WriteMsg(&msg.SCEnterGame{
		Result: 0,
	})

	c := Rooms[m.Room].Room.Open(10)
	c.Call1("join", a)
	Clients[m.Room] = c

}
