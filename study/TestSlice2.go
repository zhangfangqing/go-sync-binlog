package main

import "fmt"

func main() {

	// 声明一个长度为10的byte数组，并且赋值
	var arr = [4]byte{'a', 'b', 'c', 'd'}
	for key, value := range arr {
		fmt.Println("old arr key", key)
		fmt.Println("old arr value", value)
	}
	fmt.Println()

	// 声明2个slice，声明方式跟数组一样；区别在于不需要声明长度
	// var slice1 , slice2 []byte
	var slice1 []byte
	var slice2 []byte

	// 赋值给slice1; 注意是冒号
	slice1 = arr[0:4] // slice1=c,d(len = 3 , cap = 8)
	slice2 = arr[0:4] // slice2=d (len = 1 , cap = 7)
	// slice是引用类型，一个slice改变值；其它slice也会跟着变
	slice1[1] = 'a'

	for key, value := range arr {
		fmt.Println("new arr key", key)
		fmt.Println("new arr value", value)
	}
	fmt.Println()

	for key, value := range slice1 {
		fmt.Println("slice1 key", key)
		fmt.Println("slice1 value", value)
	}
	fmt.Println()

	for key, value := range slice2 {
		fmt.Println("slice2 key", key)
		fmt.Println("slice2 value", value)
	}
	fmt.Println()
}
