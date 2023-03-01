package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	febonaci := make([]int, n)
	febonaci[0] = 1
	febonaci[1] = 1
	for i := 2; i < n; i++ {
		febonaci[i] = febonaci[i-2] + febonaci[i-1]

	}
	for i := 0; i < n; i++ {
		fmt.Print(febonaci[i], " ")
	}
	fmt.Println()
}
