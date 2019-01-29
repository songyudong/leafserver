package internal

import (
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type RoomData struct {
	Id        int
	Agents    map[int]*gate.Agent
	Room      *chanrpc.Server
	TimeStamp int64
	Mode      int
}

var (
	RoomIdSeed = 1000
	Rooms      map[int]*RoomData
)

func init() {
	Rooms = make(map[int]*RoomData)
}

func NewRoom(mode int) int {
	log.Debug("new room seed=%v", RoomIdSeed)
	s := chanrpc.NewServer(10)
	go RoomCoroutine(s, RoomIdSeed)

	rd := new(RoomData)
	rd.Id = RoomIdSeed
	rd.Mode = mode
	RoomIdSeed++
	rd.Room = s
	rd.TimeStamp = time.Now().Unix()
	Rooms[rd.Id] = rd
	return rd.Id
}

func FindMatchRoom(mode int) int {

	for k, v := range Rooms {
		if v.Mode == mode && len(v.Agents) < mode {
			return k
		}
	}

	return NewRoom(mode)

}
