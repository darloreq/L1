package main

import (
	"context"
	"fmt"
	"time"
)

//в данном случае горутина зашершается по времени работы context

func test(ctx context.Context, exit chan bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine stopped")
			exit <- true
			return
		default:
			fmt.Println("work!")
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
	exit := make(chan bool, 1)
	go test(ctx, exit)
	<-exit
	close(exit)
}
