package main

import "fmt"

// 函数传递指针；（指针指向内存地址）
func addPoint(a *int) (int) {
	// 指针变量加上*号，等于指针指向内存的值
	*a = *a + 1
	return *a
}

func main() {
	oldVal := 3
	fmt.Println("oldVal = ", oldVal)

	// 普通变量加上&, 等于变量的内存地址
	newVal := addPoint(&oldVal)
	fmt.Println("newVal =", newVal)

}
