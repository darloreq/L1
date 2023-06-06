package main

import "fmt"

// Паттерн Адаптер (Adapter) предназначен преобразования одного класса в интерфейс, который он не поддерживает,
// к этому самому интерфейсу без изменения самогу класса. Это достигается путем создания оболочки для текущего класса,
// которая будет, в свою очередь, удовлетворяет нужному интерфейсу

type run interface {
	run()
}

// У нас есть класс runner, который удовлетворяет интерфейсу run
// Однако он не удовлетворяет интерфейсу Movement
type runner struct {
	speed int
}

func (r *runner) run() {
	fmt.Println("The runner is running at a speed of ", r.speed)
}

type Movement interface {
	Move()
}

// Создаем оболочку для класса runner, которая удовлетворяет интерфейсу Movement
type runnerAdapter struct {
	runner
}

func (r *runnerAdapter) Move() {
	r.run()
}

func main() {
	test1 := &runner{speed: 20}
	test := runnerAdapter{
		runner: *test1,
	}
	test.Move()
	test1.run()
}
