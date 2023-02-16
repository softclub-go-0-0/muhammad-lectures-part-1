package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, n float64
	fmt.Scan(&n, &a, &b)
	fmt.Printf("%.5f", ((b - a) / n))
	h := (b - a) / n
	for i := a; i <= b; i += h {
		fmt.Println(1 - math.Sin(i))
	}

}
