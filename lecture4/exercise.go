package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	evenIndex := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&evenIndex[i])
	}
	fmt.Print("Even numbers are:")
	for i := 0; i < n; i++ {
		if evenIndex[i]%2 == 0 {
			fmt.Print(evenIndex[i], " ")
		}

	}
}
