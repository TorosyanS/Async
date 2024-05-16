package times

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func CheckTimeAfter() {
	for i := 0; i < 5; i++ {
		delay := rand.Int() % 5000
		<-time.After(time.Duration(delay) * time.Millisecond)
		fmt.Printf("воркер отработал %d раз\n", i)
	}
}
