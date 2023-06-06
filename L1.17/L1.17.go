package main

import "fmt"

// бинарный поиск заключается в том, что мы берем центральный элемент массива и сравниваем его с искомым значением
// если искомый элемент, выходим
// если значение больше искомого, то повторяем пункты выше, левее среднего значения
// если значение меньше искомого, то повторяем пункты выше, правее среднего значения

func binarySearch(arr []int, n int) int {
	left := 0
	right := len(arr) - 1
	for right-left > 0 {
		middle := (right + left) / 2
		if arr[middle] == n {
			return middle
		}
		if arr[middle] < n {
			left = middle + 1
		} else if arr[middle] > n {
			right = middle - 1
		}
	}
	if arr[left] == n {
		return left
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 8, 9, 11, 15, 20, 23, 28, 29, 30}
	fmt.Println(binarySearch(arr, 32))
	fmt.Println(binarySearch(arr, 0))
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 3))
	fmt.Println(binarySearch(arr, 5))
	fmt.Println(binarySearch(arr, 8))
	fmt.Println(binarySearch(arr, 9))
	fmt.Println(binarySearch(arr, 11))
	fmt.Println(binarySearch(arr, 15))
	fmt.Println(binarySearch(arr, 20))
	fmt.Println(binarySearch(arr, 23))
	fmt.Println(binarySearch(arr, 28))
	fmt.Println(binarySearch(arr, 29))
	fmt.Println(binarySearch(arr, 30))
}
