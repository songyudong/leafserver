package msg

import (
	vector2d "server/utils"

	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})

	Processor.Register(&CSLogin{})
	Processor.Register(&CSChat{})
	Processor.Register(&CSMatch{})
	Processor.Register(&CSEnterGame{})

	Processor.Register(&SCLogin{})
	Processor.Register(&SCChat{})
	Processor.Register(&SCMatch{})
	Processor.Register(&SCEnterGame{})
	Processor.Register(&SCSpawnUnit{})
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

type SCSpawnUnit struct {
	Iid      int
	UType    int
	Pos      vector2d.Vector2D
	FaceLeft bool
	UFaction int
}

//------------------------------------
type UserData struct {
	UserId   int
	UserName string
	level    int
	money    int
}
