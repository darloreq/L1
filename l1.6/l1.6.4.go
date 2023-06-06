package main

import (
	"fmt"
	"time"
)

// В данном случае горутина останавливатеся по закрытию канала

func test2(waitChan chan interface{}) {
	for i := range waitChan {
		fmt.Println(i)
	}
	fmt.Println("goroutine stopped")
}

func main() {
	waitChan := make(chan interface{}, 1)
	go test2(waitChan)
	go func() {
		for i := 0; i < 25; i++ {
			waitChan <- i
		}
	}()
	time.Sleep(time.Second)

	close(waitChan)
	time.Sleep(time.Second)
}
