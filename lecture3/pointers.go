package main

import "fmt"

func main() {

	var a = 10
	var ptr = &a

	fmt.Println("Initial values:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

	a += 5
	fmt.Println("\nAfter incrementing a direction:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
	*ptr += 5
	fmt.Println("\nAfter incrementing a direction:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)

	secondptr := &ptr

	fmt.Println("\nAfter incrementing a direction:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
	fmt.Printf("%v\t\t%T\n", secondptr, secondptr)
	fmt.Printf("%v\t\t%T\n", *secondptr, *secondptr)
	fmt.Printf("%v\t\t%T\n", **secondptr, **secondptr)

	thirdptr := &secondptr

	fmt.Println("\nAfter incrementing a direction:")
	fmt.Printf("%v\t\t%T\n", a, a)
	fmt.Printf("%v\t\t%T\n", ptr, ptr)
	fmt.Printf("%v\t\t%T\n", *ptr, *ptr)
	fmt.Printf("%v\t\t%T\n", secondptr, secondptr)
	fmt.Printf("%v\t\t%T\n", *secondptr, *secondptr)
	fmt.Printf("%v\t\t%T\n", **secondptr, **secondptr)
	fmt.Printf("%v\t\t%T\n", secondptr, secondptr)
	fmt.Printf("%v\t\t%T\n", *thirdptr, *thirdptr)
	fmt.Printf("%v\t\t%T\n", **thirdptr, **thirdptr)
	fmt.Printf("%v\t\t%T\n", ***thirdptr, ***thirdptr)

}
