package main

import (
	"fmt"
	"sync"
)

func main() {
	// Инициализируем WaitGroup для того чтоб дождаться завершения всех потоков
	wg := sync.WaitGroup{}
	arr := []int{
		2, 4, 6, 8, 10,
	}
	// Инициализируем Mutex для конкурентной среды
	mutex := sync.Mutex{}
	for _, elem := range arr {
		// при создании горутины добавляем в счетчик wg единицу
		wg.Add(1)
		go func(elem int) {
			// по завершению функции убавляем у счетчика единицу
			defer wg.Done()
			// блочим мутекс, чтоб другой поток не смог вывести число
			mutex.Lock()
			// выводим квадрат числа
			fmt.Println(elem * elem)
			mutex.Unlock()

		}(elem)
	}
	wg.Wait()

}
