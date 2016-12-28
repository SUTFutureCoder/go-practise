package main

import  "fmt"

type Human struct {
	name    string
	age     int
	phone   string
}

type Student struct {
	Human
	school  string
	loan    float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human实现SayHI方法
func (h Human) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s \n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee 重载Human Say Hi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.Human.name, e.company, e.Human.phone)
}

//Interface Men被Human,Student,Employee实现
//因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-xxx"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-xxx"}, "Harvard", 100}
	sam  := Employee{Human{"Sam", 36, "444-222-xxx"}, "Golang INC", 1000}
	Tom  := Employee{Human{"Tom", 36, "444-222-xxx"}, "Baidu", 5000}

	//定义Men类型变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a student:")
	i.SayHi()
	i.Sing("A LILI A LILI A LI A LILI")

	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义slice Men
	fmt.Println("Let's use a slice of Men and see what happen")
	x := make([]Men, 3)
	//T这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}

}