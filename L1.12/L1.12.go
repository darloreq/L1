package main

import "fmt"

type void struct{}

// инициализируем мапу с ключем типа string, а значением типа пустой структуры, так как она весит 0 кб

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	mapped := make(map[string]void)
	for _, str := range strs {
		mapped[str] = void{}
	}
	fmt.Println(mapped)
}
