package chans

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func CheckChanWithoutBuf() {
	ch := make(chan int)

	go startWorker(ch)

	time.Sleep(2 * time.Second)

	data, isOpen := <-ch
	if isOpen {
		fmt.Println("не закрыл")
	}
	fmt.Println("Получено значение", data)

	//for data := range ch {
	//	fmt.Println("Получено значение", data)
	//}
	//data := <-ch
	//fmt.Println("Получено значение", data)
	//
	//data = <-ch
	//fmt.Println("Получено значение", data)
	//
	//data = <-ch
	//fmt.Println("Получено значение", data)
	//
	//data = <-ch
	//fmt.Println("Получено значение", data)
	//
	//data, isOpen := <-ch
	//fmt.Println("Получено значение", data)
	//if isOpen {
	//	fmt.Println("Канал все еще открыт")
	//} else {
	//	fmt.Println("Канал уже закрыт")
	//}
	//
	//data, isOpen = <-ch
	//fmt.Println("Получено значение", data)
	//if isOpen {
	//	fmt.Println("Канал все еще открыт")
	//} else {
	//	fmt.Println("Канал уже закрыт")
	//}
}

func startWorker(ch chan int) {
	for i := 0; i < 10; i++ {
		select {
		case ch <- rand.Int():
			fmt.Println("Положили в канал новое значение")
		default:
			fmt.Println("Никто не ждет чтения")
			time.Sleep(time.Second)
		}
	}
	close(ch)
	fmt.Println("канал закрыт")
}
