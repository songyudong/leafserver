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

const (
	UF_Blue int = iota
	UF_Red
)

type Unit struct {
	IsDeleted bool
	Iid       int
	UType     int
	Pos       utils.Vector2D
	FaceLeft  bool
	UFaction  int
}

type Player struct {
	Iid     int
	Agent   *gate.Agent
	Faction int
	UIid    int
}

type Collider struct {
}
