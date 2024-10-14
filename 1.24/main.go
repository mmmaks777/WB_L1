package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func Distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

func main() {
	a := NewPoint(1, 1)
	b := NewPoint(2, 2)

	fmt.Println(Distance(a, b))
}
