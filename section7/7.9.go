package main

import (
	"fmt"
	"os"
)

func main() {
	str := "section7/FutureCoder.txt"
	file, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break;
		}
		os.Stdout.Write(buf[:n])
	}
}
