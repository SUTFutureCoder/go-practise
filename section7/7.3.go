package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
)

type Server struct {
	ServerName  string
	ServerIP    string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	var s Serverslice
	f, err := os.Open("section7/servers.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	str, err := ioutil.ReadAll(f)
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers[0].ServerIP)

}
