package main

import "fmt"

func add1(a *int) *int {
	*a = *a + 1
	return a
}

func add2(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3

	fmt.Printf("x = %d\n", x)

	x1 := add1(&x)

	fmt.Printf("x1 = %d\n", *x1)
	fmt.Printf("x  = %d\n", x)

	x2 := add2(&x)
	fmt.Printf("x2 = %d\n", x2)
	fmt.Printf("x  = %d\n", *&x)


}
