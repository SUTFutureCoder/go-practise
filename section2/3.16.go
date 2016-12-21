package main

//并发
import (
	"fmt"
	"runtime"
)

func say(s string){
	runtime.GOMAXPROCS(16)
	for i := 0; i < 5; i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main(){
	go say("HELLO")
	say("WORLD")
}
