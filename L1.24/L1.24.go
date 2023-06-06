package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int // инкапсулируем x и y
	y int
}

// NewPoint - конструктор
func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(p2 *Point) (distance float64) {
	// формула нахждения дистанции между точками
	distance = math.Sqrt(math.Pow(float64(p2.x-p.x), 2) + math.Pow(float64(p2.y-p.y), 2))
	return
}

func main() {
	p1 := NewPoint(-111, -12)
	p2 := NewPoint(32, 13)
	fmt.Println(p1.Distance(p2))
}
