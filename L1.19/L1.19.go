package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	// переводим нашу строку в руну для того, что не потерять значения юникодов
	arr := []rune(os.Args[1])
	// далее двумя указателями расположенными слева и справа строки еняем данные местами
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(string(arr))
}
