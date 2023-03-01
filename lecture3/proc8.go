package main

import "fmt"

func main() {
	var k, d int
	fmt.Scan(&k)
	for i := 1; i <= 2; i++ {
		fmt.Scan(&d)
		AddRightDigit(&k, &d)
		fmt.Println(k)
	}
}
func AddRightDigit(k, d *int) {
	*k = *k*10 + *d
}
