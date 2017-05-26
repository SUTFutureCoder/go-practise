package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	i := 0
	for {
		fmt.Fprintf(conn, "hello server\n")
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%#v\n", data)
		fmt.Println(i)
		i++
	}


}
