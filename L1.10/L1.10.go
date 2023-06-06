package main

import "fmt"

// функция окргуления
func toRes(val float32) int {
	return int(val) / 10 * 10
}

func main() {
	temper := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	mapped := make(map[int][]float32)
	for _, i := range temper {
		key := toRes(i)
		mapped[key] = append(mapped[key], i)
	}
	fmt.Println(mapped)
}
