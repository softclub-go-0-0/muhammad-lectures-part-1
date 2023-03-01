package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	if n > 2 {
		formNum := make([]int, n)
		formNum[0] = a
		formNum[1] = b
		formNum[2] = a + b
		for i := 3; i < n; i++ {
			s:=i

			formNum[i] += formNum[i-1]
		}
		for i := 0; i < n; i++ {
			fmt.Print(formNum[i], " ")
		}
	}
}
