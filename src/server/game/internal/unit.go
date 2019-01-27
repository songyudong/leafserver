package internal

import (
	vector2d "server/utils"

	"github.com/name5566/leaf/gate"
)

const (
	UT_None    int = iota // value --> 0
	UT_Hero               // value --> 1
	UT_Agent              // value --> 2
	UT_Monster            // value --> 3
)

const (
	UF_Blue int = iota
	UF_Red
)

type Unit struct {
	IsDeleted bool
	Iid       int
	UType     int
	Pos       vector2d.Vector2D
	FaceLeft  bool
	UFaction  int
}

type Player struct {
	Iid     int
	Agent   *gate.Agent
	Faction int
}

type Collider struct {
}
