package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"
	fmt.Printf("%s\n", m["Hello"])

}
