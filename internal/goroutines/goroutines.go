package goroutines

import (
	"fmt"
	"time"
)

var counter int

func doReq() {
	time.Sleep(time.Millisecond)
	counter++
}

func CheckWithoutGoroutines() {
	for i := 0; i < 2000; i++ {
		doReq()
	}

	fmt.Println(counter)
}

func CheckWithGoroutines() {
	for i := 0; i < 2000; i++ {
		go doReq()
	}

	fmt.Println(counter)
}
