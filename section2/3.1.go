package main

import "fmt"

func main() {
	integer := 6
	switch integer {
	case 4:
		fmt.Printf("The integer was <=4")
		fallthrough
	case 5:
		fmt.Printf("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Printf("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Printf("The integer was <= 7")
		fallthrough
	default:
		fmt.Printf("DEFAULT")
	}
}
