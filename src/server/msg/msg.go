package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&CSLogin{})
	Processor.Register(&CSChat{})

	Processor.Register(&SCLogin{})
	Processor.Register(&SCChat{})

	Processor.Register(&UserData{})
}

type Hello struct {
	Name string
}

type CSLogin struct {
	UserName string
	Password string
}

type SCLogin struct {
	ErrorCode int
	UserId    int
}

type CSChat struct {
	Content string
}

type SCChat struct {
	UserId   int
	UserName string
	Content  string
}

type UserData struct {
	UserId   int
	UserName string
}
