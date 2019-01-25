package internal

import (
	"fmt"
	"server/base"

	"github.com/name5566/leaf/db/mongodb"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	conn     *mongodb.DialContext
)

type Module struct {
	*module.Skeleton
}

type DBUser struct {
	UserId   int
	UserName string
	Age      int
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	log.Debug("111111111111111111111111")

	c, err := mongodb.Dial("localhost", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer c.Close()
	//s := c.Ref()

	conn = c

	//c.EnsureCounter("banana", "user", "username")
	//id1, err := c.NextSeq("banana", "user", "username")
	//id2, err := c.NextSeq("banana", "user", "username")

	/*u := DBUser{
		id1,
		"song",
		25,
	}
	u1 := DBUser{
		id2,
		"guan",
		26,
	}

	col := s.DB("banana").C("user")

	col.DropCollection()
	col.Insert(&u)
	col.Insert(&u1)

	update := DBUser{
		id1,
		"songyudong",
		25,
	}
	col.Update(bson.M{"userid": id1}, update)
	query := col.Find(bson.M{"age": bson.M{"$eq": 25}})
	us := []DBUser{}
	query.All(&us)
	log.Debug("query result %v", us)

	q2 := col.Find(nil)
	us2 := []DBUser{}
	q2.All(&us2)
	log.Debug("query result2 %v", us2)

	u3 := DBUser{}
	q2.One(&u3)
	log.Debug("query result3 %v", u3)

	q3 := col.Find(bson.M{"username": "songyudong"}).Select(bson.M{"_id": 1})
	var r interface{}
	q3.One(&r)
	log.Debug("query result4 %v", r)
	q4 := col.Find(bson.M{"_id": r.(bson.M)["_id"]})
	u4 := DBUser{}
	q4.One(&u4)
	log.Debug("query result5 %v", u4)

	col.Update(bson.M{"username": "songyudong"}, bson.M{"$set": bson.M{"age": 33}})

	*/
}

func (m *Module) OnDestroy() {

}

func (m *Module) Login(username string) {

}
