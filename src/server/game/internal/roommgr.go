package internal

import (
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/gate"
)

type RoomData struct {
	Id        int
	Agents    map[int]*gate.Agent
	Room      *chanrpc.Server
	TimeStamp int64
}

var (
	RoomIdSeed = 1000
	Rooms      map[int]*RoomData
)

func init() {
	Rooms = make(map[int]*RoomData)
}

func NewRoom() int {
	s := chanrpc.NewServer(10)
	go RoomCoroutine(s)

	rd := new(RoomData)
	rd.Id = RoomIdSeed
	RoomIdSeed++
	rd.Room = s
	rd.TimeStamp = time.Now().Unix()
	Rooms[rd.Id] = rd
	return rd.Id
}
