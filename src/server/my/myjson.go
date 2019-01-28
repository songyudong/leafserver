package mychan

import (
	"encoding/json"

	"io/ioutil"

	"github.com/name5566/leaf/log"
)

type ColliderData struct {
	R []float64
}

type ColliderRoot struct {
	Nodes []ColliderData
}

type JsonStruct struct {
}

func ExampleJson() {

	JsonParse := &JsonStruct{}
	v := ColliderRoot{}

	JsonParse.Load("gamedata/colliderRoot.json", &v)

	log.Debug("json string : %v", v)
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
