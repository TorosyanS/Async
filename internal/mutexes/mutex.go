package mutexes

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func doReq(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	mu.Lock()
	counter++
	mu.Unlock()
}

func CheckMutex() {
	wg := sync.WaitGroup{}
	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go doReq(&wg)
	}

	wg.Wait()
	fmt.Println(counter)
}
