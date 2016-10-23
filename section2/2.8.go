package main

import "fmt"

func main() {
	//	var numbers map[string]int

	//另一种map声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Printf("three %d\n", numbers["three"])
	fmt.Printf("len %d\n", len(numbers))

}
