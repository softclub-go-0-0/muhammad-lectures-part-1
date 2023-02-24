package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, amean, gmean float64

	fmt.Scan(&a, &b, &c, &d)
	mean(a, b, &amean, &gmean)
	fmt.Println(amean, " ", gmean)
	mean(a, c, &amean, &gmean)
	fmt.Println(amean, " ", gmean)
	mean(a, d, &amean, &gmean)
	fmt.Println(amean, " ", gmean)
}
func mean(x, y float64, amean, gmean *float64) {
	*amean = (x + y) / 2
	*gmean = math.Sqrt(x * y)
}
