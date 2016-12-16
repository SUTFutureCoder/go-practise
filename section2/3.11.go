package main

import  "fmt"

type Human struct {
	name string
	age  int
	telephone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h Human) sayHi() {
	fmt.Println("Hi, I am ", h.name, " You can call me on", h.telephone)
}

func (s Student) sayHi() {
	fmt.Println("Hi, I am a student named" , s.name, " Study in ", s.school)
}

func (e Employee) sayHi() {
	fmt.Println("Hi, I am a employee named", e.name, " Employeed in", e.company)
}

func main() {
	student  := Student{Human{"LinXingchen", 22, "15101669791"}, "Shenyang Universion Of Technology"}
	employee := Employee{Human{"Xingchen", 22, "15101669791"}, "Baidu"}

	student.sayHi()
	employee.sayHi()
}