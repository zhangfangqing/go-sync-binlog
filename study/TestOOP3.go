package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	score int
}

type Emploo struct {
	Human
	salar float32
}

func (h Human) sayHi() {
	fmt.Println("human sayhi")
}

// Student 重写sayHi方法
func (e Emploo) sayHi() {
	fmt.Println("emploo sayHi")
}

func main() {

	stu := Student{Human{"fangqing", 29}, 90}
	stu.sayHi()

	emploo := Emploo{Human{"mude", 54}, 150.23}
	emploo.sayHi()
}
