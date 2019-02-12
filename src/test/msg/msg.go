package msg

import (
	
	"test/message"
	"github.com/name5566/leaf/network/sproto"
)

var Processor = sproto.NewProcessor()

func init() {
	Processor.Register(&message.CSChat{})
	Processor.Register(&message.CSDrop{})
	Processor.Register(&message.CSEnterGame{})
	Processor.Register(&message.CSFloat{})
	Processor.Register(&message.CSLogin{})
	Processor.Register(&message.CSMatch{})
	Processor.Register(&message.CSMove{})
	Processor.Register(&message.CSStop{})
	Processor.Register(&message.Header{})
	Processor.Register(&message.Hello{})
	
	Processor.Register(&message.SCBlowCancel{})
	Processor.Register(&message.SCBlowStart{})
	Processor.Register(&message.SCBlowSuccess{})
	Processor.Register(&message.SCBurst{})
	Processor.Register(&message.SCChat{})
	Processor.Register(&message.SCDrop{})
	Processor.Register(&message.SCEnterGame{})
	Processor.Register(&message.SCFire{})

	Processor.Register(&message.SCFloat{})
	Processor.Register(&message.SCGameStart{})
	Processor.Register(&message.SCGameState{})
	Processor.Register(&message.SCLogin{})
	Processor.Register(&message.SCMatch{})
	Processor.Register(&message.SCMove{})
	Processor.Register(&message.SCSpawnUnit{})

	Processor.Register(&message.SCStop{})
	Processor.Register(&message.UnitState{})
	Processor.Register(&message.UserData{})
	Processor.Register(&message.Vector2D{})
	
	
}


