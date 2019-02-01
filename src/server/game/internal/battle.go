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
	FrameNumber   int
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
	b.CurTime = 0
	b.FrameNumber = 0
	for _, v := range b.Players {
		(*v.Agent).WriteMsg(&msg.SCGameStart{
			TimeStamp: b.StartTime,
		})
	}
	b.FrameNumber++
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
		u.Pos.Y = 110
	} else {
		u.FaceLeft = true
		u.Pos.X = 600
		u.Pos.Y = 110
	}
	u.LastPos = u.Pos
	u.Ballons = 1
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

	//log.Debug("battle.update")
	b.CurTime = float64(time.Now().UnixNano())/NANO - b.StartTime

	//log.Debug("---------------------------------------------")
	for {
		stop := false
		select {
		case data := <-b.Server.ChanCall:
			b.Server.Exec(data)
			//log.Debug("111111111111111111111111111111111111111111")
			break
		default:
			stop = true
			break
		}

		if stop {
			break
		}
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
	//log.Debug("update logic")
	for _, v := range b.Units {
		OldPos := v.Pos
		v.LastPos = v.Pos
		PosH := v.Pos
		PosV := v.Pos
		NewX := v.Pos.X
		NewY := v.Pos.Y
		if v.Moving {
			if v.FaceLeft {
				NewX -= float64(c_xVel) * delta
			} else {
				NewX += float64(c_xVel) * delta
			}

			PosH.X = NewX
		}

		if v.Floating {
			NewY += float64(c_yFloatVel) * delta
		} else {
			NewY -= float64(c_yDropVel) * delta
		}
		PosV.Y = NewY

		v.Pos.X = ClampScreen(v.Pos.X)
		v.Pos = utils.Vector2D{X: NewX, Y: NewY}
		nr := v.GetCollider()
		if !IntersectWithWorld(nr) {
			v.ClampScreen()
			continue
		}
		/*v.Pos = PosH
		nr = v.GetCollider()
		if !IntersectWithWorld(nr) {

			if v.Moving {
				continue
			}

		}

		v.Pos = PosV
		nr = v.GetCollider()
		if !IntersectWithWorld(nr) {

			continue
		}*/

		v.Pos = OldPos
		safePos := OldPos

		STEP := 10
		if v.Moving {
			for i := 0; i < STEP; i++ {
				if v.FaceLeft {
					v.Pos.X -= float64(c_xVel) * delta / float64(STEP)
				} else {
					v.Pos.X += float64(c_xVel) * delta / float64(STEP)
				}

				v.Pos.X = ClampScreen(v.Pos.X)

				nr = v.GetCollider()
				if !IntersectWithWorld(nr) {
					safePos = v.Pos
				} else {
					v.Pos = safePos

					break
				}
			}
		}

		for i := 0; i < STEP; i++ {
			if v.Floating {
				v.Pos.Y += float64(c_yFloatVel) * delta / float64(STEP)
			} else {
				v.Pos.Y -= float64(c_yDropVel) * delta / float64(STEP)
			}
			nr = v.GetCollider()
			if !IntersectWithWorld(nr) {
				safePos = v.Pos
			} else {
				v.Pos = safePos

				break
			}
		}

		v.ClampScreen()

	}

	for _, u1 := range b.Units {
		col1 := u1.GetFoot()
		for _, u2 := range b.Units {
			if u2.Iid == u1.Iid {
				continue
			}
			if u2.Ballons > 0 {
				bal2 := u2.GetBallon()
				if utils.IsIntersect(col1, bal2) {
					u1.Score++
					u2.Burst()
					for _, v := range b.Players {
						(*v.Agent).WriteMsg(&msg.SCBurst{
							Iid: u2.Iid,
						})
					}
				}
			}

		}
	}

	for _, v := range b.Units {
		if v.Blowing {
			if v.Pos != v.LastPos {
				log.Debug("blow cancel")
				v.BlowCancel()
				for _, v := range b.Players {
					(*v.Agent).WriteMsg(&msg.SCBlowCancel{
						Iid: v.Iid,
					})
				}
				continue
			}
			v.BlowTimer -= delta
			log.Debug("blow timer = %v, delta = %v", v.BlowTimer, delta)
			if v.BlowTimer <= 0 {
				v.BlowSuccess()
				for _, vv := range b.Players {
					(*vv.Agent).WriteMsg(&msg.SCBlowSuccess{
						Iid: v.Iid,
					})
				}
			}
		}
	}
	b.SendState()
	b.FrameNumber++
}

func (b *Battle) SendState() {
	//log.Debug("send state")
	State := msg.SCGameState{}
	State.CurTime = b.CurTime
	State.FrameNumber = b.FrameNumber
	for _, v := range b.Units {
		State.UnitStates = append(State.UnitStates, msg.UnitState{
			Iid:      v.Iid,
			Pos:      v.Pos,
			FaceLeft: v.FaceLeft,
			Moving:   v.Moving,
			Floating: v.Floating,
			Blowing:  v.Blowing,
		})
	}

	for _, v := range b.Players {
		(*v.Agent).WriteMsg(&State)
	}
}

func IntersectWithWorld(r *utils.Rect) bool {

	for _, v := range colliders {
		if utils.IsIntersect(r, v) {
			//log.Debug("intersect")
			return true
		}

	}

	return false
}

func ClampScreen(x float64) float64 {
	if x < 0 {
		return x + c_zoneWidth
	} else if x >= c_zoneWidth {
		return x - c_zoneWidth
	}

	return x
}
