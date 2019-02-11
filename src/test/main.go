package main

import (
	"test/conf"
	_ "test/db"
	_ "test/gamedata"
	
	
	"test/gate"
	"test/login"
	

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
		
		gate.Module,
		login.Module,
		//db.Module,
		
	)

}
