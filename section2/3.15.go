package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.14
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	//修改必须这样写
	p := reflect.ValueOf(&x)
	e := p.Elem()
	e.SetFloat(6.66)
	fmt.Println("v-value:", e)
}
