package times

import (
	"fmt"
	"time"
)

func CheckTicker() {
	ticker := time.Tick(1 * time.Second)

	for {
		<-ticker
		fmt.Println("Воркер отработал")
	}
}
