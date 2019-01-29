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
	handler(&msg.CSMove{}, handlerMove)
	handler(&msg.CSStop{}, handlerStop)
	handler(&msg.CSFloat{}, handlerFloat)
	handler(&msg.CSDrop{}, handlerDrop)
	handler(&msg.CSFire{}, handlerFire)

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
	log.Debug("c2s match mode %v", m.Mode)

	RoomId := 0
	if m.Mode == 1 {
		RoomId = NewRoom(m.Mode)
	} else {
		RoomId = FindMatchRoom(m.Mode)
	}

	a.WriteMsg(&msg.SCMatch{
		Result: 0,
		Room:   RoomId,
	})
}

func handlerEnterGame(args []interface{}) {
	m := args[0].(*msg.CSEnterGame)
	a := args[1].(gate.Agent)
	log.Debug("c2s match room %v", m.Room)

	a.WriteMsg(&msg.SCEnterGame{
		Result: 0,
	})

	c := Rooms[m.Room].Room.Open(10)
	c.Call1("join", a)
	Clients[m.Room] = c

	ud := a.UserData().(*mongodbmgr.DBUser)
	ud.Room = m.Room
	a.SetUserData(ud)
}

func handlerMove(args []interface{}) {
	m := args[0].(*msg.CSMove)
	a := args[1].(gate.Agent)
	log.Debug("user %v c2s move left=%v", a.UserData().(*mongodbmgr.DBUser).UserId, m.Left)
	ud := a.UserData().(*mongodbmgr.DBUser)
	c := Clients[ud.Room]
	log.Debug("call move %v", c)
	c.Call1("move", a, m.Left)
}

func handlerStop(args []interface{}) {
	m := args[0].(*msg.CSStop)
	a := args[1].(gate.Agent)
	log.Debug("user %v c2s stop %v", a.UserData().(*mongodbmgr.DBUser).UserId, m)
	ud := a.UserData().(*mongodbmgr.DBUser)
	c := Clients[ud.Room]
	log.Debug("call stop %v", c)
	c.Call1("stop", a)
}

func handlerFloat(args []interface{}) {
	m := args[0].(*msg.CSFloat)
	a := args[1].(gate.Agent)
	log.Debug("user %v c2s float %v", a.UserData().(*mongodbmgr.DBUser).UserId, m)
	ud := a.UserData().(*mongodbmgr.DBUser)
	c := Clients[ud.Room]
	log.Debug("call float %v", c)
	c.Call1("float", a)
}

func handlerDrop(args []interface{}) {
	m := args[0].(*msg.CSDrop)
	a := args[1].(gate.Agent)
	log.Debug("user %v c2s drop %v", a.UserData().(*mongodbmgr.DBUser).UserId, m)
	ud := a.UserData().(*mongodbmgr.DBUser)
	c := Clients[ud.Room]
	log.Debug("call drop %v", c)
	c.Call1("drop", a)
}

func handlerFire(args []interface{}) {
	m := args[0].(*msg.CSFire)
	a := args[1].(gate.Agent)
	log.Debug("user %v c2s fire %v", a.UserData().(*mongodbmgr.DBUser).UserId, m)
}
