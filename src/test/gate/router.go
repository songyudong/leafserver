package gate

import (
	
	"test/login"
	"test/message"
	"test/msg"
)

func init() {
	msg.Processor.SetRouter(&message.Hello{}, login.ChanRPC)
	
}
