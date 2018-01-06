package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
}

func main() {
	// 声明一个结构体
	var p1 person
	// 赋值
	p1.name = "fangqing"
	p1.age = 29
	fmt.Println("fangqing name = ", p1.name)

	// 声明方式2；按照顺序提供初始化值
	p2 := person{"xiaomin", 28}
	fmt.Println("xiaomin name =", p2.name)

	//声明方式3；按照field:value的方式初始化，这样可以任意顺序
	p3 := person{age: 24, name: "tome"}
	fmt.Println("tom name = ", p3.name)
}
