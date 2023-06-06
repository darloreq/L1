package main

import (
	"fmt"
	"sync"
)

// стрктура содежит переменную типа инт - является счетчиком
// и мутекс для конкурентности
type Counter struct {
	count int
	mutex sync.Mutex
}

//метод инкремента значения, в нем мутекс лочится, что другие потоки не могли его менять
func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.count++
}

//метод получения данных , в нем мутекс лочится, что другие потоки не могли поменять значения до его получения
func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func work(c *Counter, finish chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int, wg *sync.WaitGroup, c *Counter) {
			fmt.Println("Worker", id, "start")
			defer wg.Done()
			c.Inc()
			fmt.Println("Worker", id, "finish")
		}(i+1, &wg, c)
	}
	wg.Wait()
	finish <- struct{}{}
	close(finish)
}

func main() {
	co := &Counter{}
	finish := make(chan struct{}, 1)
	go work(co, finish)
	<-finish
	fmt.Println(co.GetValue())
}
