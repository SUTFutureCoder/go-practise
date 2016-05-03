package main

import "fmt"
import "errors"

func main() {
	err := errors.New("emit macho dwarft: elf header corrupted")
	if err != nil {
		fmt.Printf("%v", err)
	}
}
