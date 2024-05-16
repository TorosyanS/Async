package atomics

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter atomic.Int32

func doReq() {
	time.Sleep(100 * time.Millisecond)
	counter.Add(1)
}

func CheckAtomics() {
	for i := 0; i < 2000; i++ {
		go doReq()
	}

	fmt.Println(counter.Load())
}
