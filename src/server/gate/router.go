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
}
