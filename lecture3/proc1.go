package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("A: ")
	fmt.Scan(&a)
	powerA(&b, a)
	fmt.Println("B:", b)
}
func powerA(b *float64, a float64) {
	*b = a * a * a
}
