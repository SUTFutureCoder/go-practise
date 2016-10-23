package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [66]int{1, 2, 3}
	for index := 0; index < 3; index++ {
		fmt.Printf("a-%d:%d\n", index, a[index])
		fmt.Printf("b-%d:%d\n", index, b[index])
		fmt.Printf("c-%d:%d\n", index, c[index])
	}
	d := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	for index := 0; index < 2; index++ {
		for index2 := 0; index2 < 4; index2++ {
			fmt.Printf("d-%d\n", d[index][index2])
		}
	}
}
