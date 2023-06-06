package main

import (
	"fmt"
	"strings"
)

type UniqueStr map[rune]bool

func UniqueString(str string) bool {
	mapped := make(UniqueStr)
	// переводим всю строку в нижний регистр
	str = strings.ToLower(str)
	for _, i := range str {
		// если в мапе присутствует элемент из строки, то возвр. false
		if _, ok := mapped[i]; ok {
			return false
		}
		// иначе добавляем в мапу
		mapped[i] = true
	}
	return true
}

func main() {
	// строки с повторениеб без и пустая
	str := "asdfA"
	str1 := "qweRTY"
	str2 := ""
	fmt.Println(UniqueString(str))
	fmt.Println(UniqueString(str1))
	fmt.Println(UniqueString(str2))
}
