package contexts

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

type ctxKey string

const myKey ctxKey = "myKey"

const i = 500

func CheckContext() {
	ticker := time.Tick(1 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	notifyCtx, _ := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Истек контекст")
			break LOOP
		case <-notifyCtx.Done():
			fmt.Println("Кто-то остановил процесс")
			break LOOP
		case <-ticker:
			fmt.Println("Воркер отработал")
		}
	}
}

func CheckContextWithValue() {
	ctx := context.WithValue(context.Background(), myKey, i)

	checkCtxValue(ctx)
}

func checkCtxValue(ctx context.Context) {
	value := ctx.Value(myKey)
	ctxValue, ok := value.(int)
	if !ok {
		panic("incorrect type")
	}

	fmt.Printf("в контексте лежал %d\n", ctxValue)
}
