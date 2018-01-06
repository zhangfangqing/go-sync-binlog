package main

import "fmt"

type Rectangle struct {
	height float32
	weight float32
}

// 函数传递一个结构体做为参数；跟方式不一样；方法在func添加一个接受者；方式属于某个接受者
func area(r Rectangle) (float32) {
	return r.height * r.weight
}

func main() {
	r := Rectangle{12, 5}
	fmt.Println(area(r))

	r2 := Rectangle{height: 2, weight: 3}
	fmt.Println(area(r2))

}
