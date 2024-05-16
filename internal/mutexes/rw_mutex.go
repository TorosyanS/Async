package mutexes

import (
	"fmt"
	"sync"
	"time"
)

var (
	dictionary map[int]int
	rwMu       sync.RWMutex
)

func doRequest(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	time.Sleep(time.Millisecond)
	isWriter := number%1000 == 0
	if isWriter {
		rwMu.Lock()
		dictionary[number] = number
		rwMu.Unlock()
	} else {
		rwMu.RLock()
		_, _ = dictionary[number]
		rwMu.RUnlock()
	}
}

func CheckRWMutex() {
	now := time.Now()

	defer func() {
		fmt.Printf("Программа выполнилась за %d миллисекунд", time.Since(now).Milliseconds())
	}()
	wg := sync.WaitGroup{}
	dictionary = make(map[int]int, 100)
	for i := 0; i < 2000000; i++ {
		wg.Add(1)
		go doRequest(&wg, i)
	}

	wg.Wait()
	fmt.Println(len(dictionary))
}
