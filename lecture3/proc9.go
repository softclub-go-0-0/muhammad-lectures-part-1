package main

import (
	"fmt"
	"math"
)

func main() {
	var k, d int
	fmt.Scan(&k)
	for i := 1; i <= 2; i++ {
		fmt.Scan(&d)
		AddLeftDigit(&k, &d)
		fmt.Println(k)
	}
}
func AddLeftDigit(k, d *int) {
	var cnt int
	s := *k
	for i := s; i >= 1; i /= 10 {
		*k = 0
		cnt++
	}
	var p float64
	p = math.Pow(10, float64(cnt))
	*k = int(p)**d + s
}
