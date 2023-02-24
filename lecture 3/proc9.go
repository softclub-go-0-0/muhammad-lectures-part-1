package main

import "fmt"

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
	*k = *k + *d*10
}
