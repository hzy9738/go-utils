package main

import (
	queue2 "github.com/hzy9738/common-utils/queue"
	"log"
	"sync"
	"time"
)

func main() {
	queue := queue2.Map{
		C:   make(map[string]*queue2.Entry),
		Rmx: &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			val := queue.Rd("key", time.Second*15)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			queue.Out("key", val)
		}(i)
	}

	time.Sleep(time.Second * 30)
}
