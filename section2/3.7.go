package main

import "fmt"

//声明一个新的类型
type person struct {
	name string
	age  int
}

func older(p1, p2 person) (person, int){
	if p1.age > p2.age {
		//比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person

	//赋值初始化
	tom.name, tom.age = "Tom", 18

	//两个字段都写清楚的初始化
	van := person{"Van", 44}

	//按照struct定义顺序初始化值
	billy := person{name: "Billy", age: 47}

	tv_older, tv_diff := older(tom, van)
	tb_older, tb_diff := older(tom, billy)
	vb_older, vb_diff := older(van, billy)

	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, van.name, tv_older.name, tv_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", tom.name, billy.name, tb_older.name, tb_diff)
	fmt.Printf("Of %s and %s, %s is older by %d years\n", van.name, billy.name, vb_older.name, vb_diff)


}
