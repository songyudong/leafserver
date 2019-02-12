package gamedata

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/songyudong/leaf/log"
	"github.com/songyudong/leaf/recordfile"
)

var (
	PhyData *ColliderRoot
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".txt"
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}

	return rf
}

type ColliderData struct {
	R []float64
}

type ColliderRoot struct {
	Nodes []ColliderData
}

type JsonStruct struct {
}

func init() {

	JsonParse := &JsonStruct{}
	v := ColliderRoot{}

	JsonParse.Load("gamedata/colliderRoot.json", &v)

	log.Debug("json string : %v", v)

	PhyData = &v
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Debug("not find file")
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Debug("unmarshal failed")
		return
	}
}
