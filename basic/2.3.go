package main

import "fmt"

const PI = 3.1415926
const i = 100000000
const MaxThread = 10
const prefix = "astaxie_"

func main() {
	var enabled, disabled = true, false
	fmt.Printf("%v %v %v %v %v %v", PI, i, MaxThread, prefix, enabled, disabled)
}
