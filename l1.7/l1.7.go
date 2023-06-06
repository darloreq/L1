package main

import (
	"fmt"
	"log"
	"sync"
)

// конкурентная запись данных в мап осуществляется с помощью мутексов, так как мапдля запиине потокобезопасный

func main() {
	mapped := make(map[int]interface{})
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 26; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			mapped[id] = 0
			fmt.Println(id, "append in map")
		}(i + 1)
	}
	wg.Wait()
	for i, t := range mapped {
		log.Println(i, t)
	}
}
