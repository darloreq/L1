package main

import "fmt"

func main() {
	var (
		a int8
		b int8
	)
	a = 127
	b = 5
	fmt.Println(a, b)
	//a, b = b, a
	//fmt.Println(a, b)
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}
