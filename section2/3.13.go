package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name    string
	age     int
	phone   string
}
//通过这个方法Human实现了 fmt.Stringer
func (h Human) String() string {
	return h.name + "-" + strconv.Itoa(h.age) + " years - ☏ " + h.phone
}

func main(){
	Bob := Human{"Bob", 39, "000-7777-xxx"}
	fmt.Println("This human is :", Bob)
}
