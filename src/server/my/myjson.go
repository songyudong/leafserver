package mychan

import (
	"io/ioutil"
	"os"

	"github.com/name5566/leaf/log"
)

func ExampleJson() {
	fi, err := os.Open("gamedata/colliderRoot.json")
	if err != nil {
		log.Debug("not find file")
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	log.Debug("content : %v", fd)
}
