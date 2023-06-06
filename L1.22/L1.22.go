package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(36e17)
	f := big.NewInt(13e16)
	fmt.Println("Деление", a.Div(a, f))
	fmt.Println("Умножение", a.Mul(a, f))
	fmt.Println("Сложение", a.Add(a, f))
	fmt.Println("Вычитание", a.Sub(f, a))
}
