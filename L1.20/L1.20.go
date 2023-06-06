package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	newStr := ""
	// разбиваем нашу строку на массив строк, деление происход по пробелу
	strs := strings.Split(str, " ")
	// после чего добавляем все елементы с конца в строку
	for i := len(strs) - 1; i >= 0; i-- {
		newStr += fmt.Sprint(strs[i])
		if i != 0 {
			newStr += " "
		}
	}
	return newStr
}

func main() {
	str := "qwe asd zxc rty фыв bvn"
	fmt.Println(str)
	str = reverseString(str)
	fmt.Println(str)
}
