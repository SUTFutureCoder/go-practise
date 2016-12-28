package main

import (
	"fmt"
	"os"
)

func main() {
	userFile    := "section7/FutureCoder.txt"
	fout, err   := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\n");
		fout.Write([]byte("Just a test!\n"));
	}
}
