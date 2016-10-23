package main

import "fmt"

func main() {
	s := "Hello world"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s", s2)

	s3 := "Hello"
	s3 = "c" + s3[1:]
	fmt.Printf("%s\n", s3)
}
