package main

import "fmt"

func main() {

	// 声明数组，方式1
	var arr [10]int
	// 赋值操作
	arr[0] = 11
	arr[1] = 12

	// 声明数组，方式2
	arr2 := [3]int{1, 2, 3}
	arr3 := [10]int{111, 222, 333}
	arr4 := [...]int{444, 555, 666}
	for key, value := range arr2 {
		fmt.Println("arr2 key", key)
		fmt.Println("arr2 value", value)
	}
	for key, value := range arr3 {
		fmt.Println("arr3 key", key)
		fmt.Println("arr3 value", value)
	}
	for key, value := range arr4 {
		fmt.Println("arr4 key", key)
		fmt.Println("arr4 value", value)
	}

	// 二维数组声明
	doubleArr1 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	dounleArr2 := [2][4]int{{11, 22, 33, 44}, {55, 66, 77, 88}}
	for key, value := range doubleArr1 {
		fmt.Println("doubleArr1 key", key)
		fmt.Println("doubleArr1 value", value)
	}
	for key, value := range dounleArr2 {
		fmt.Println("doubleArr1 key", key)
		fmt.Println("doubleArr1 value", value)
	}

}
