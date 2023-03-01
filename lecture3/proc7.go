package main

import (
	"fmt"
)

func invDigits(n *int) {
	var s int
	s = *n
	*n = 0
	for i := s; i >= 1; i /= 10 {

		*n = *n*10 + i%10
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	invDigits(&n)
	fmt.Println(n)
}
