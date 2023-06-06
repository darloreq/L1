package main

import "log"

func main() {
	var man Action
	// у структуры Action есть методы структуры Human
	man.SetName("Tom")
	log.Println(man.GetName())
}
