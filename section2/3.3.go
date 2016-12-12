package main

import "fmt"

func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("xPLUSy %d\n", xPLUSy)
	fmt.Printf("xTIMESy %d\n", xTIMESy)

}
