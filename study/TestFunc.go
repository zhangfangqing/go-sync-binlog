package main

import "fmt"

//返回一个参数
func add(a, b int) int {
	return a + b
}

// 返回1个参数
func add1(a, b int) (int) {
	return a + b
}

// 返回1个参数
func add2(a, b int) (sum int) {
	return a + b
}

// 返回2个值
func addMuti(a, b int) (sum, muti int) {
	return a + b, a * b
}

// 返回2个值
func addMuti2(a, b int) (string, int) {
	return "ok", a * b
}

// 可变参数
func addParamChange(arg ...int) (sum int) {
	ret := 0
	for _, value := range arg {
		ret = ret + value
	}
	return ret
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add1(2, 3))

	sum, mutil := addMuti(3, 5)
	fmt.Println(sum)
	fmt.Println(mutil)

	fmt.Println(addParamChange(1, 2, 3, 4, 5))
	fmt.Println(addParamChange(1, 2))
}
