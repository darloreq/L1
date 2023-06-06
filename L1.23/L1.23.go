package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 0
	_, err := fmt.Scan(&i)
	if err != nil || i <= 0 || i > len(arr) {
		fmt.Println("Введи нормально...")
		return
	}
	i -= 1
	// сдвигаем все элементы начиная с i-ого на один левее
	copy(arr[i:], arr[i+1:])
	// убираем последний элемент
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
