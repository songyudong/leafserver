package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.CSChat{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSMatch{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSEnterGame{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSMove{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSStop{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSFloat{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSDrop{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CSFire{}, game.ChanRPC)
}
