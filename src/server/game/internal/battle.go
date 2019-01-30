package internal

import (
	mongodbmgr "server/db"
	"server/gamedata"
	"server/msg"
	"server/utils"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
)

var colliders []*utils.Rect

type Battle struct {
	Server        *chanrpc.Server
	Units         map[int]*Unit
	UnitsAdd      map[int]*Unit
	Players       map[int]*Player
	UnitIidSeed   int
	PlayerIidSeed int
	CurSlot       int
	StartTime     float64
	CurTime       float64
}

func init() {
	for k, v := range gamedata.PhyData.Nodes {
		log.Debug("%v, %v", k, v)
		r := utils.Rect{
			X:      v.R[0],
			Y:      v.R[1],
			Width:  v.R[2],
			Height: v.R[3],
		}
		colliders = append(colliders, &r)
	}
}

func (b *Battle) init() {
	b.UnitIidSeed = 2000
	b.PlayerIidSeed = 1000
	b.Units = make(map[int]*Unit)
	b.UnitsAdd = make(map[int]*Unit)
	b.Players = make(map[int]*Player)
}

func (b *Battle) firstFrame() {
	log.Debug("first frame")
	b.SpawnHeroes()
	b.updateMembers()

	b.StartTime = float64(time.Now().UnixNano()) / NANO
	b.CurTime = b.StartTime
}

func (b *Battle) addUnit(u *Unit) {
	b.UnitsAdd[u.Iid] = u
}

func (b *Battle) addPlayer(p *Player) {
	b.Players[p.Iid] = p
}

func (b *Battle) updateMembers() {
	for k, v := range b.UnitsAdd {
		b.Units[k] = v
	}
	b.UnitsAdd = make(map[int]*Unit)
}

func (b *Battle) SpawnHero(p *Player) *Unit {
	u := new(Unit)
	u.UType = UT_Hero
	u.UFaction = p.Faction
	u.Iid = b.UnitIidSeed
	b.UnitIidSeed++
	if u.UFaction == UF_Blue {
		u.FaceLeft = false
		u.Pos.X = 100
		u.Pos.Y = 100
	} else {
		u.FaceLeft = true
		u.Pos.X = 600
		u.Pos.Y = 100
	}
	p.UIid = u.Iid

	ud := (*p.Agent).UserData().(*mongodbmgr.DBUser)
	log.Debug("exe join userid=%v", ud.UserId)

	b.addUnit(u)

	for _, v := range b.Players {
		(*v.Agent).WriteMsg(&msg.SCSpawnUnit{
			Iid:      u.Iid,
			UType:    u.UType,
			Pos:      u.Pos,
			FaceLeft: u.FaceLeft,
			UFaction: u.UFaction,
			UserId:   p.Iid,
		})
	}
	return u
}

func (b *Battle) SpawnPlayer(agent *gate.Agent, faction int) *Player {
	p := new(Player)
	ud := (*agent).UserData().(*mongodbmgr.DBUser)
	p.Iid = ud.UserId
	p.Agent = agent
	p.Faction = faction
	return p
}

func (b *Battle) SpawnHeroes() {
	for _, v := range b.Players {
		b.SpawnHero(v)
	}
}

func (b *Battle) Update(delta float64) {
	select {
	case data := <-b.Server.ChanCall:
		b.Server.Exec(data)
	default:
	}

	b.UpdateLogic(delta)

	d := timer.NewDispatcher(10)
	duration := time.Duration(66 * MSEC)
	d.AfterFunc(duration, func() {
		b.Update(0.066)
	})
	(<-d.ChanTimer).Cb()
}

func (b *Battle) UpdateLogic(delta float64) {
	for _, v := range b.Units {
		OldPos := v.Pos
		PosH := v.Pos
		PosV := v.Pos
		NewX := v.Pos.X
		NewY := v.Pos.Y
		if v.Moving {
			if v.FaceLeft {
				NewX -= float64(xVel) * delta
			} else {
				NewX += float64(xVel) * delta
			}

			PosH.X = NewX
		}

		if v.Floating {
			NewY += float64(yFloatVel) * delta
		} else {
			NewY -= float64(yDropVel) * delta
		}
		PosV.Y = NewY

		v.Pos = utils.Vector2D{X: NewX, Y: NewY}
		nr := v.GetRect()
		if !IntersectWithWorld(nr) {
			continue
		}
		v.Pos = PosH
		nr = v.GetRect()
		if !IntersectWithWorld(nr) {
			continue
		}

		v.Pos = PosH
		nr = v.GetRect()
		if !IntersectWithWorld(nr) {
			continue
		}

		v.Pos = OldPos
	}
}

func IntersectWithWorld(r *utils.Rect) bool {

	for _, v := range colliders {
		if utils.IsIntersect(r, v) {
			return true
		}

	}

	return false
}
