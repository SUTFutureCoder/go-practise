package main

import "fmt"

func main() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	a := ar[2:5]
	//b := ar[3:5]

	for index := 0; index < 3; index++ {
		fmt.Printf("%c\n", a[index])
	}
}
