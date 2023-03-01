package main

import (
	"fmt"
)

func main() {
	var n int
	var a, d float64
	fmt.Scan(&n, &a, &d)
	aProgression := make([]float64, n)
	aProgression[0] = a
	for i := 1; i < n; i++ {
		aProgression[i] = aProgression[i-1] * d
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%.2f\t", aProgression[i])
	}
	fmt.Println()
}


