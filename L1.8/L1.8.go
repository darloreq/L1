package main

import "fmt"

func main() {
	var number int64 = 100
	var bit int64 = 1
	one := false
	switch one {
	case true:
		// с помощью дизъюнкции устанавливаем значение байта единицей
		// пример 101 | 010 = 111
		number = number | (1 << bit)
	default:
		// c помощью сброса бита устанавливаем значение бита нулем
		// пример 111 &^ 100 = 011
		number = number &^ (1 << bit)
	}
	fmt.Println("result is - ", number)
}
