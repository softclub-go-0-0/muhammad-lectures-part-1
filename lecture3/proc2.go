package main

import "fmt"

func main() {
	var a, b, c, d float64
	fmt.Print("A: ")
	fmt.Scan(&a)
	powerA234(&b, &c, &d, a)
	fmt.Println("B:", b)
	fmt.Println("C:", c)
	fmt.Println("D:", d)

}
func powerA234(b, c, d *float64, a float64) {
	*b = a * a
	*c = a * a * a
	*d = a * a * a * a
}
