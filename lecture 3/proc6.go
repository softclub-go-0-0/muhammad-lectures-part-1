package main

import "fmt"

func main() {
	var k, c, s int
	fmt.Scan(&k)
	digitcount(k, &s, &c)
	fmt.Println(c, " ", s)
}
func digitcount(k int, c, s *int) {
	for i := k; i >= 1; i /= 10 {

		*s++
		*c += i % 10

	}
}
