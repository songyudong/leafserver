package msg

import (
	
	"test/message"
	"github.com/name5566/leaf/network/sproto"
)

var Processor = sproto.NewProcessor()

func init() {
	Processor.Register(&message.Hello{})
	Processor.Register(&message.SCLogin{})
	
}


