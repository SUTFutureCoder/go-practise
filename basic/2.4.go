package main

import "fmt"

func main() {
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiou"
	frenchHello := "Bonjour"

	fmt.Printf("%v %v %v %v %v\n", no, yes, maybe, japaneseHello, frenchHello)

	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%v", s2)

	s3 := "Hello "
	s4 := "World"
	s3 += s4
	fmt.Printf("\n%v\n", s3)

	s5 := "c" + s3[1:]
	fmt.Printf("%v\n", s5)

	m := `Hello
		World`

	fmt.Printf("%v\n", m)
}
