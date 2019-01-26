package internal

import "github.com/name5566/leaf/gate"

type Room struct {
	Id     int
	Agents map[int]*gate.Agent

	TimeStamp int64
}
