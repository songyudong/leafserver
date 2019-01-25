package internal

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	skeleton.RegisterChanRPC("login", rpcLogin)
}

func rpcLogin(args []interface{}) {
	log.Debug("rpcLogin")

	username := args[0].(string)
	log.Debug(username)

	s := conn.Ref()
	col := s.DB("banana").C("user")
	query := col.Find(bson.M{"username": username})
	var result interface{}
	query.One(&result)
	log.Debug("%v", result)
	if result == nil {
		log.Debug("not exist")
	} else {
		log.Debug("find succses")
	}
}
