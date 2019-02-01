package internal

import (
	mongodbmgr "server/db"
	"server/msg"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
)

func RoomCoroutine(s *chanrpc.Server, roomId int, mode int) {
	log.Debug("room coroutine roomid=%v", roomId)
	battle := new(Battle)
	battle.init()

	battle.Server = s

	waitJoin := mode

	s.Register("join", func(args []interface{}) interface{} {
		log.Debug("player join")
		a := args[0].(gate.Agent)

		f := UF_Blue
		if battle.CurSlot%2 == 0 {
			f = UF_Blue
		} else {
			f = UF_Red
		}
		battle.CurSlot++
		np := battle.SpawnPlayer(&a, f)
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
		u.Moving = true
		u.FaceLeft = left
		//log.Debug("unit=%v", u)
		//log.Debug("player move left=%v", left)
		/*for _, v := range battle.Players {
			(*v.Agent).WriteMsg(&msg.SCMove{
				Iid:  u.Iid,
				Left: left,
			})
		}*/
		return 0
	})

	s.Register("stop", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		u.Moving = false
		//log.Debug("player stop")
		/*for _, v := range battle.Players {
			(*v.Agent).WriteMsg(&msg.SCStop{
				Iid: u.Iid,
			})
		}*/
		return 0
	})

	s.Register("float", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		//log.Debug("ballons count = %v", u.Ballons)
		if u.CanFloat() {
			//log.Debug("--------------------------------")
			u.Floating = true
		} else if !u.Blowing && u.Stand() {
			u.Blowing = true
			u.BlowTimer = 3
			for _, v := range battle.Players {
				(*v.Agent).WriteMsg(&msg.SCBlowStart{
					Iid: u.Iid,
				})
			}
		}
		//log.Debug("player float")

		return 0
	})

	s.Register("drop", func(args []interface{}) interface{} {
		a := args[0].(gate.Agent)
		p := battle.Players[a.UserData().(*mongodbmgr.DBUser).UserId]
		u := battle.Units[p.UIid]
		u.Floating = false
		if u.Blowing {
			u.BlowCancel()
			for _, v := range battle.Players {
				(*v.Agent).WriteMsg(&msg.SCBlowCancel{
					Iid: u.Iid,
				})
			}
		}
		//log.Debug("player drop")
		/*for _, v := range battle.Players {
			(*v.Agent).WriteMsg(&msg.SCDrop{
				Iid: u.Iid,
			})
		}*/
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

	d := timer.NewDispatcher(10)
	duration := time.Duration(66 * MSEC)
	d.AfterFunc(duration, func() {
		battle.Update(0.066)
	})
	(<-d.ChanTimer).Cb()
}
