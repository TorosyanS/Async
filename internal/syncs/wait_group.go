package syncs

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counter atomic.Int32

func doReq(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	counter.Add(1)
}

func CheckWaitGroup() {
	now := time.Now()
	defer func() {
		fmt.Printf("Работа заняла %d миллисекунды\n", time.Since(now).Milliseconds())
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go doReq(&wg)
	}

	wg.Wait()
	fmt.Println(counter.Load())
}
