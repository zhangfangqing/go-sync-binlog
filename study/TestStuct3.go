package main

import "fmt"

// slice 类型
type Skills []string

// 结构体
type Humen struct {
	name string
	age  int
}

// 结构体
type Man struct {
	Humen
	Skills
	specia string
}

func main() {
	fangqing := Man{Humen{"fangqing", 29}, []string{"weiqi", "xiangqi"}, "program"}
	fmt.Println("old Man= ", fangqing)

	//修改skill
	fangqing.Skills = []string{"weiqi"}
	fmt.Println("new Man= ", fangqing)
}
