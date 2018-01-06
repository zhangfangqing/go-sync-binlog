package main

import "fmt"

// 长方形结构体
type ChangFangXing struct {
	height float32
	weight float32
}

// 圆形结构体
type YuanXing struct {
	a float32
}

// 为长方形添加方法
func (c ChangFangXing) area() (float32) {
	return c.height * c.weight
}

// 为圆形添加方法
func (r YuanXing) area() (float32) {
	return r.a * r.a
}

func main() {
	c := ChangFangXing{height: 12, weight: 5}
	fmt.Println(c.area())

	y := YuanXing{a: 3.14}
	fmt.Println(y.area())
}
