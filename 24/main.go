package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func main() {
	firstPoint := NewPoint(0, 0)
	secondPoint := NewPoint(1, 1)
	fmt.Println("Расстояние:", GetDistance(firstPoint, secondPoint))
}

//Конструктор структуры Point
func NewPoint(x, y float64) *Point {
	newSt := &Point{
		X: x,
		Y: y,
	}

	return newSt
}
func GetDistance(f, s *Point) float64 {
	return math.Sqrt((s.X-f.X)*(s.X-f.X) + (s.Y-f.Y)*(s.Y-f.Y))
}
