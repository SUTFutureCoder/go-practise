package main

import "fmt"

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0]) //获取数据
	fmt.Printf("The last element is %d\n", arr[9])
}
