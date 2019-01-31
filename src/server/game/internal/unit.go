package internal

import (
	"server/utils"

	"github.com/name5566/leaf/gate"
)

const (
	UT_None    int = iota // value --> 0
	UT_Hero               // value --> 1
	UT_Agent              // value --> 2
	UT_Monster            // value --> 3
)

const NANO float64 = 1000000000
const MSEC int64 = 1000000

const (
	UF_Blue int = iota
	UF_Red
)

const xVel int = 100
const yFloatVel int = 60
const yDropVel int = 160

type Unit struct {
	IsDeleted bool
	Iid       int
	UType     int
	Pos       utils.Vector2D
	FaceLeft  bool
	UFaction  int
	Moving    bool
	Floating  bool
	Ballons   int
}

func (u *Unit) GetRect() *utils.Rect {
	r := utils.Rect{
		X:      u.Pos.X - 20/2,
		Y:      u.Pos.Y,
		Width:  20,
		Height: 58,
	}
	return &r
}

type Player struct {
	Iid     int
	Agent   *gate.Agent
	Faction int
	UIid    int
}

type Collider struct {
}
