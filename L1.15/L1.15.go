package main

import (
	"fmt"
)

func createHugeString(len int) (str string) {
	for i := 0; i < len; i++ {
		str += "g"
	}
	return str
}

func createHugeStringNew(len int) (str string) {
	for i := 0; i < len; i++ {
		str += "ш"
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	v1 := createHugeStringNew(1 << 10)
	fmt.Println("len(v) - ", len(v))
	fmt.Println("len(v1) - ", len(v1))
	juststring := v[:100]
	juststring1 := v1[:100]
	// проблемой данного кода является отсутствие обработки unicode сомволов
	// так как в string unicode занимает 2 байта, то срез строки приведет к тому,
	// что половина символов потеряется вовсе
	// также проблемой является то, что срез указывает на саму строку и сборщик мусора не сможет почистить
	// полную строку по той причине, что на неё ссылается срез, поэтому данные нужно копировать в новый slice
	newjuststring := make([]rune, 0, 100)
	for i := 0; i < 100; i++ {
		newjuststring = append(newjuststring, rune([]rune(v1)[i]))
	}
	fmt.Println("len(juststring) - ", len(juststring))
	fmt.Println("len(juststring1) - ", len(juststring1))
	fmt.Println("len(newjuststring) - ", len(newjuststring))
	fmt.Println(juststring)
	fmt.Println(juststring1)
	fmt.Println(string(newjuststring))

}

func main() {
	someFunc()
}
