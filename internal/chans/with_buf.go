package chans

import (
	"fmt"
	"time"
)

func CheckChanWithBuf() {
	ch := make(chan int, 9)

	go startWorker(ch)

	time.Sleep(2 * time.Second)

	data, isOpen := <-ch
	if isOpen {
		fmt.Println("не закрыл")
	}
	fmt.Println("Получено значение", data)
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
	//
	//for data = range ch {
	//	fmt.Println("Получено значение", data)
	//}
}
