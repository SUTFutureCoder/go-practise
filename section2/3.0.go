package main

import "fmt"

func main() {
	map1 := make(map[string]string)
	map1["string"] = "string1"
	map1["hello"] = "hello world"
	map1["baidu"] = "baidu"
	for k, v := range map1 {
		fmt.Printf("key is %s\n", k)
		fmt.Printf("value is %s\n", v)
	}
}
