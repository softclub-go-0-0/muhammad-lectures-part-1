package main

import (
	"fmt"
	"math"
)

func main() {
	var a, p, s float64
	fmt.Scan(&a)
	trianglePS(&p, &s, a)
	fmt.Println(p, " ", s)
}

func trianglePS(p, s *float64, a float64) {
	*p = 3 * a
	*s = a * a * math.Sqrt(3) / 4
}
