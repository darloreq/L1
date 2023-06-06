package main

type Human struct {
	name string
	age  int8
}

func (man *Human) SetName(name string) {
	man.name = name
}

func (man *Human) SetAge(age int8) {
	man.age = age
}

func (man *Human) GetName() string {
	return man.name
}

func (man *Human) GetAge() int8 {
	return man.age
}

// Наследование методов в go реализовано с помощью композиции
type Action struct {
	Human
}
