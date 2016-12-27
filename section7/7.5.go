package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName    string
	Age         int
}

func main() {
	t       := template.New("fieldname example")
	t, _    = t.Parse("hello {{.}} {{.UserName}}")
	p := Person{UserName: "FutureCoder", Age: 17}
	t.Execute(os.Stdout, p)
}
