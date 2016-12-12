package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x := 3
	y := 4
	z := 5

	max_xy := max(x, y)
	max_xz := max(x, z)

	fmt.Printf("max xy%d\n", max_xy)
	fmt.Printf("max xz%d\n", max_xz)
}
