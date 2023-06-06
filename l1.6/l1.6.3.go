package main

import (
	"log"
	"sync"
	"time"
)

// в данном случае горутина останавливается по своему циклу жизни

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			log.Println(id, "goroutine run")
			time.Sleep(time.Second)
			log.Println(id, "goroutine a closed")
		}(i + 1)
	}
	wg.Wait()
}
