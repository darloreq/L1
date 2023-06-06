package main

import "fmt"

func defineType(data interface{}) {
	// получаем тип данных полученного значеня
	switch v := data.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case bool:
		fmt.Println("bool", v)
	case chan int:
		fmt.Println("chan int", v)
	default:
		fmt.Println(" hz", v)
	}
}

func main() {
	var (
		a int
		b string
		c bool
		d chan int
	)
	defineType(a)
	defineType(b)
	defineType(c)
	defineType(d)
}
