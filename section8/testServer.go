package main

import (
	"net"
	"log"
	"fmt"
	"bufio"
)

func handleConnection(conn net.Conn){
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("get client data error: ", err)
	}
	fmt.Printf("%#v\n", data)
	fmt.Fprintf(conn, "hello client\n")
	//conn.Close()
}

func main() {
	link, err := net.Listen("tcp", ":10086")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := link.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}
