package time

import (
	"sync"
	"time"
)

// 要求sync.WaitGroup支持timeout功能
// 如果timeout到了超时时间返回true
// 如果WaitGroup自然结束返回false

func WaitTimeOut(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan bool, 1)
	go time.AfterFunc(timeout, func() {
		ch <- true
	})
	go func() {
		wg.Wait()
		ch <- false
	}()
	return <-ch
}
