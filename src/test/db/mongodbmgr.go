package mongodbmgr

import (
	"fmt"

	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
)

type DBUser struct {
	UserId   int
	UserName string
	Age      int
	Room     int
}

var dialContext = new(mongodb.DialContext)

func init() {
	Connect()
}

func Connect() {
	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer c.Close()
	// index
	s := c.Ref()
	col := s.DB("banana").C("user")

	count, _ := col.Count()
	if count == 0 {
		c.EnsureCounter("banana", "user", "username")
	}
	dialContext = c

}

//
func Login(username string) (*DBUser, int) {
	s := dialContext.Ref()
	col := s.DB("banana").C("user")
	query := col.Find(bson.M{"username": username})
	//var result interface{}
	ud := DBUser{}
	//query.One(&result)
	query.One(&ud)
	log.Debug("%v", ud)
	if ud.UserId == 0 {
		log.Debug("not exist")
		id, _ := dialContext.NextSeq("banana", "user", "username")

		u := DBUser{
			id,
			username,
			25,
			0,
		}
		col.Insert(&u)
		return &u, 0
	} else {
		log.Debug("find succses %v", ud)

		return &ud, 1
	}
}
