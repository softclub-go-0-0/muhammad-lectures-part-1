package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) GetQuarter() int {
	//need to implement

	return 0
}

type Line struct {
	start, finish Point
}

func (l *Line) GetLength() float64 {
	return math.Sqrt((math.Pow(l.finish.x-l.start.x, 2) + math.Pow(l.finish.y-l.start.y, 2)))
}

type Square struct {
	Side Line
}

func (s *Square) GetPerimeter() float64 {
	return 4 * s.Side.GetLength()
}
func (s *Square) GetArea() float64 {
	return math.Pow(s.Side.GetLength(), 2)
}
func main() {
	square := Square{
		Side: Line{
			start: Point{
				x: 4,
				y: 4,
			},
			finish: Point{
				x: 0,
				y: 4,
			},
		},
	}
	fmt.Println(square)
	fmt.Println(square.GetPerimeter())
	fmt.Println(square.GetArea())
}
