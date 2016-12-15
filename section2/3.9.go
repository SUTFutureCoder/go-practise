package main

import "fmt"

type Skills []string

type Human struct {
	name string
	age  int
	weight int
}

type Student struct {
	Human  //匿名字段，struct
	Skills //匿名字段，自定义的类型string slice
	int    //内置类型作为匿名字段
	speciality string
}

func main() {
	//初始化学生Jane
	jane := Student{Human:Human{"Jane", 35, 100}, speciality:"CS"} //注意格式统一
	fmt.Printf("Her name is %s\n", jane.Human.name)
	fmt.Printf("Her age  is %s\n", jane.Human.age)
	fmt.Printf("Her weight is %s\n", jane.Human.weight)
	fmt.Printf("Her speciality is %s\n", jane.speciality)
	//修改skill字段
	jane.Skills = []string{"anatomy"}
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skill now is \n", jane.Skills)
	//修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}
