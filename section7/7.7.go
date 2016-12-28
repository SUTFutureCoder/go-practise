package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("FutureCoder", 0777)
	os.MkdirAll("FutureCoder/test1/test2", 0777)
	err := os.Remove("FutureCoder")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("FutureCoder")
}
