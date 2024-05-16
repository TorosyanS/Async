package selects

import (
	"fmt"
	"time"
)

func CheckSelect() {
	slowTicker := time.NewTicker(1 * time.Second)
	fastTicker := time.Tick(1 * time.Second)
	//ch := make(chan int)
	//count := 0

	//LOOP:
	//for {
	select {
	//case ch <- 15:
	//	fmt.Println("все получилось")
	case <-slowTicker.C:
		fmt.Println("отработал медленный тикер")
		//count++
		//if count == 5 {
		//	slowTicker.Stop()
		//	break LOOP
		//}
	case <-fastTicker:
		fmt.Println("отработал быстрый тикер")
		//
		//default:
		//	fmt.Println("произошла катастрофа, буфер забился")
		//}
	}
	//}
}
