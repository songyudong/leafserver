package internal

import (
	mongodbmgr "server/db"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type Battle struct {
	Units         map[int]*Unit
	UnitsAdd      map[int]*Unit
	Players       map[int]*Player
	UnitIidSeed   int
	PlayerIidSeed int
	CurSlot       int
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
	/*(*p.Agent).WriteMsg(&msg.SCSpawnUnit{
		Iid:      u.Iid,
		UType:    u.UType,
		Pos:      u.Pos,
		FaceLeft: u.FaceLeft,
		UFaction: u.UFaction,
		UserId:   p.Iid,
	})*/

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
