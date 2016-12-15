package main

import "fmt"

type Human struct {
	name string
	age  int
	weight int
}

type Student struct {
	Human //匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	//初始化一个学生
	mark := Student{Human{"Mark", 15, 120}, "Computer Science"}

	//访问相应的字段
	fmt.Printf("His name is%s\n", mark.name)
	fmt.Printf("His age  is%d\n", mark.age)
	fmt.Printf("His age  is%d\n", mark.Human.age)
	fmt.Printf("His weight is%d\n", mark.weight)
	fmt.Printf("His speciality is%s\n", mark.speciality)
	//修改对应字段信息
	mark.speciality = "AI"
	fmt.Printf("%s changed his spec iality to %s\n", mark.name, mark.speciality)
}
