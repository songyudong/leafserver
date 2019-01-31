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

const c_xVel int = 100
const c_yFloatVel int = 60
const c_yDropVel int = 160
const c_zoneWidth float64 = 960
const c_zoneHeight float64 = 540
const c_heroWidth float64 = 20
const c_heroHeight float64 = 58
const c_ballonYOff float64 = 55
const c_ballonWidth float64 = 55
const c_ballonHeight float64 = 53

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

func (u *Unit) GetCollider() *utils.Rect {
	r := utils.Rect{
		X:      u.Pos.X - c_heroWidth/2,
		Y:      u.Pos.Y,
		Width:  c_heroWidth,
		Height: c_heroHeight,
	}
	return &r
}

func (u *Unit) GetBallon() *utils.Rect {
	r := utils.Rect{
		X:      u.Pos.X - c_ballonWidth/2,
		Y:      u.Pos.Y + c_ballonYOff,
		Width:  c_ballonWidth,
		Height: c_ballonHeight,
	}
	return &r
}

func (u *Unit) ClampScreen() {
	//log.Debug("clamp screen x=%v, y=%v", u.Pos.X, u.Pos.Y)
	if u.Pos.X < 0 {
		u.Pos.X = u.Pos.X + c_zoneWidth
	} else if u.Pos.X >= c_zoneWidth {
		u.Pos.X -= c_zoneWidth
	}
	if u.Pos.Y >= c_zoneHeight-c_heroHeight {
		//log.Debug("detect y bigger than max")
		u.Pos.Y = c_zoneHeight - c_heroHeight
	}
}

type Player struct {
	Iid     int
	Agent   *gate.Agent
	Faction int
	UIid    int
}

type Collider struct {
}
