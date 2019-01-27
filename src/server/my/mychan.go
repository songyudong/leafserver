package mychan

import (
	"fmt"
	"sync"
	"time"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
)

func init() {
	//Example()
	//ExampleData()
}

func Dosomething() {
	log.Debug("chanrpc_test.....")
}

func Example() {
	s := chanrpc.NewServer(10)

	var wg sync.WaitGroup
	wg.Add(1)

	// goroutine 1
	go func() {
		s.Register("f0", func(args []interface{}) {
			log.Debug("exe f0")
		})

		s.Register("f1", func(args []interface{}) interface{} {
			log.Debug("exe f1")
			return 1
		})

		s.Register("fn", func(args []interface{}) []interface{} {
			log.Debug("exe fn")
			return []interface{}{1, 2, 3}
		})

		s.Register("add", func(args []interface{}) interface{} {
			log.Debug("exe add")
			n1 := args[0].(int)
			n2 := args[1].(int)
			return n1 + n2
		})

		wg.Done()

		log.Debug("++++++++++++++++++++++++++++")
		for {
			log.Debug("*************************")
			s.Exec(<-s.ChanCall)
		}
	}()

	wg.Wait()
	wg.Add(1)

	// goroutine 2
	go func() {
		c := s.Open(10)

		// sync
		log.Debug("f0")
		err := c.Call0("f0")
		if err != nil {
			fmt.Println(err)
		}

		log.Debug("f1")
		r1, err := c.Call1("f1")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r1)
		}

		log.Debug("fn")
		rn, err := c.CallN("fn")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rn[0], rn[1], rn[2])
		}

		log.Debug("add")
		ra, err := c.Call1("add", 1, 2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ra)
		}

		// asyn
		log.Debug("a f0")
		c.AsynCall("f0", func(err error) {
			if err != nil {
				fmt.Println(err)
			}
		})

		log.Debug("a f1")
		c.AsynCall("f1", func(ret interface{}, err error) {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ret)
			}
		})

		log.Debug("a fn")
		c.AsynCall("fn", func(ret []interface{}, err error) {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ret[0], ret[1], ret[2])
			}
		})

		log.Debug("a add")
		c.AsynCall("add", 1, 2, func(ret interface{}, err error) {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(ret)
			}
		})

		time.Sleep(10 * time.Second)
		log.Debug("-------------------")
		c.Cb(<-c.ChanAsynRet)
		log.Debug("111111111111111111")
		c.Cb(<-c.ChanAsynRet)
		log.Debug("22222222222222222")
		c.Cb(<-c.ChanAsynRet)
		log.Debug("333333333333333")
		c.Cb(<-c.ChanAsynRet)

		log.Debug("44444444444444")
		// go
		s.Go("f0")

		wg.Done()
	}()

	wg.Wait()

	// Output:
	// 1
	// 1 2 3
	// 3
	// 1
	// 1 2 3
	// 3
}
