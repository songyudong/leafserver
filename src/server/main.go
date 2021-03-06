package main

import (
	"server/conf"
	_ "server/db"
	_ "server/gamedata"
	"server/game"
	
	"server/gate"
	"server/login"
	"server/match"

	//mychan "server/my"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	//mychan.ExampleJson()
	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
		//db.Module,
		match.Module,
	)

}
