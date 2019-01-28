package internal

import (
	mongodbmgr "server/db"
	"server/msg"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func RoomCoroutine(s *chanrpc.Server) {
	battle := new(Battle)
	battle.init()

	waitJoin := 1

	s.Register("join", func(args []interface{}) interface{} {
		log.Debug("player join")
		a := args[0].(gate.Agent)
		np := battle.SpawnPlayer(&a, UF_Blue) // todo
		battle.addPlayer(np)

		waitJoin--
		return 0
	})

	s.Register("move", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		left := args[1].(bool)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		log.Debug("player=%v", p)
		u := battle.Units[p.UIid]
		log.Debug("unit=%v", u)
		log.Debug("player move left=%v", left)
		a.WriteMsg(&msg.SCMove{
			Iid:  u.Iid,
			Left: left,
		})
		return 0
	})

	s.Register("stop", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		log.Debug("player stop")
		a.WriteMsg(&msg.SCStop{
			Iid: u.Iid,
		})
		return 0
	})

	s.Register("float", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		log.Debug("player float")
		a.WriteMsg(&msg.SCFloat{
			Iid: u.Iid,
		})
		return 0
	})

	s.Register("drop", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		log.Debug("player drop")
		a.WriteMsg(&msg.SCDrop{
			Iid: u.Iid,
		})
		return 0
	})

	for {

		s.Exec(<-s.ChanCall)
		if waitJoin == 0 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	battle.firstFrame()
	log.Debug("first frame finished")
	for {

		s.Exec(<-s.ChanCall)
		log.Debug("room chan recieve data")
	}
}
