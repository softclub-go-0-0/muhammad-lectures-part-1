package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	evenNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		evenNumbers[i] = math.Pow(2, (i + 1))
	}
	fmt.Println(evenNumbers)
}
