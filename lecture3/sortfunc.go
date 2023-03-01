package main

import "fmt"

func sortinc(a, b, c *float64) {
	var a2, b2, c2 float64
	a2 = *a
	b2 = *b
	c2 = *c
	if a2 >= b2 && a2 >= c2 && b2 >= c2 {
		*a, *b, *c = *c, *b, *a
	} else if a2 >= b2 && a2 >= c2 && c2 >= b2 {
		*a, *b, *c = *b, *c, *a
	} else if b2 >= a2 && b2 >= c2 && a2 >= c2 {
		*a, *b, *c = *c, *a, *b
	} else if b2 >= a2 && b2 >= c2 && c2 >= a2 {
		*a, *b, *c = *a, *c, *b
	} else if c2 >= a2 && c2 >= b2 && a2 >= b2 {
		*a, *b, *c = *b, *a, *c
	} else if c2 >= a2 && c2 >= b2 && b2 >= a2 {
		*a, *b, *c = *a, *b, *c
	}
}

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	sortinc(&a, &b, &c)
	fmt.Println(a, b, c)
}
