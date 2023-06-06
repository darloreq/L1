package main

import (
	"fmt"
)

// записывает все данные из слафса в мапу
func sliceToMap(set []int) map[int]int {
	mapped := make(map[int]int)
	for _, i := range set {
		mapped[i]++
	}
	return mapped
}

func main() {
	setOne := []int{1, 10, 3, 6, 4, 2, 22, 11, 55, 100, 0, 2}
	setTwo := []int{1, 6, 4, 2, 22, 7, 99, 0}
	result := make([]int, 0)
	//записываем все данные из первого слайса в мапу
	mapped := sliceToMap(setOne)
	// пробегаемся по второму слайсу , ищем совпадения в мапе, если есть, то добавляем в результатирующий слайс
	for _, i := range setTwo {
		_, ok := mapped[i]
		if ok == true {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}
