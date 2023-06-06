package main

import (
	"context"
	"fmt"
	"time"
)

// В данном случае горутина завершается ручным вызовом канала внутри контекста

func test1(ctx context.Context, exit chan bool) {
	for {
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("goroutine stopped")
			exit <- true
			return
		default:
			fmt.Println("work1!")
		}
	}
}

func main() {
	ctx := context.Background()
	exit := make(chan bool, 1)
	ctx, cancel := context.WithCancel(ctx)
	go test1(ctx, exit)
	time.Sleep(time.Second)
	cancel()
	<-exit
	close(exit)
}
