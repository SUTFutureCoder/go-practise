package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":10086"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	//i := 0
	for {
		//fmt.Println(i)
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		fmt.Println(daytime)
		_, err   = conn.Write([]byte(daytime))
		if err != nil {
			fmt.Println(err)
		}

		conn.Close()
		//fmt.Println(i)
		//i++
	}
}

func checkError(e error){
	if e != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", e.Error())
		os.Exit(1)
	}
}