package msg

import (
	"server/utils"

	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})

	Processor.Register(&CSLogin{})
	Processor.Register(&CSChat{})
	Processor.Register(&CSMatch{})
	Processor.Register(&CSEnterGame{})
	Processor.Register(&CSMove{})
	Processor.Register(&CSStop{})
	Processor.Register(&CSFloat{})
	Processor.Register(&CSDrop{})
	Processor.Register(&CSFire{})

	Processor.Register(&SCLogin{})
	Processor.Register(&SCChat{})
	Processor.Register(&SCMatch{})
	Processor.Register(&SCEnterGame{})
	Processor.Register(&SCGameStart{})
	Processor.Register(&SCSpawnUnit{})
	Processor.Register(&SCGameState{})
	Processor.Register(&SCMove{})
	Processor.Register(&SCStop{})
	Processor.Register(&SCFloat{})
	Processor.Register(&SCDrop{})
	Processor.Register(&SCFire{})
	Processor.Register(&UserData{})
}

type Hello struct {
	Name string
}

//------------------------------------
type CSLogin struct {
	UserName string
	Password string
}

type CSChat struct {
	Content string
}

type CSMatch struct {
	Mode int
}

type CSEnterGame struct {
	Room int
}

type CSMove struct {
	Left bool
}

type CSStop struct {
}

type CSFloat struct {
}

type CSDrop struct {
}

type CSFire struct {
}

//------------------------------------
type SCLogin struct {
	ErrorCode int
	UserId    int
}

type SCChat struct {
	UserId   int
	UserName string
	Content  string
}

type SCMatch struct {
	Result int
	Room   int
}

type SCEnterGame struct {
	Result int
}

type SCGameStart struct {
	TimeStamp float64
}

type SCSpawnUnit struct {
	Iid      int
	UType    int
	Pos      utils.Vector2D
	FaceLeft bool
	UFaction int
	UserId   int
}

type SCGameState struct {
	CurTime     float64
	FrameNumber int
	UnitStates  []UnitState
}

type SCMove struct {
	Iid  int
	Left bool
}

type SCStop struct {
	Iid int
}

type SCFloat struct {
	Iid int
}

type SCDrop struct {
	Iid int
}

type SCFire struct {
	Iid int
}

//------------------------------------
type UserData struct {
	UserId   int
	UserName string
	level    int
	money    int
}

type UnitState struct {
	Iid      int
	Pos      utils.Vector2D
	FaceLeft bool
	Moving   bool
	Floating bool
}
