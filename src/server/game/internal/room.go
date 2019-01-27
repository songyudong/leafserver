package internal

import (
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

	for {

		s.Exec(<-s.ChanCall)
		if waitJoin == 0 {
			break
		}
	}

	time.Sleep(5 * time.Second)
	battle.firstFrame()

	for {

		s.Exec(<-s.ChanCall)

	}
}
