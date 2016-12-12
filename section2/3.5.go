package main

import (
	"fmt"
	"os"
)

func testFile() bool{
	//f, err := os.Open("/home/kphp");
	f, err := os.Open("/home/key.php");
	defer f.Close();
	if err != nil {
		fmt.Printf("error")
		return false
	}
	b := make([]byte, 10000)
	n, _ := f.Read(b)
	fmt.Printf(string(b[:n]))
	return true
}

func testDefer() bool{
	for i := 1; i < 5; i++{
		defer fmt.Printf("%d\n", i)
	}
	return true
}

func main() {
	testFile();

	testDefer();
}
