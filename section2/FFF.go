package main

import "fmt"

func main() {
	FFF := "FFF"
	for {
		go fmt.Printf("%s123\n", FFF);
		go fmt.Printf("%s456\n", FFF);
	}
}
