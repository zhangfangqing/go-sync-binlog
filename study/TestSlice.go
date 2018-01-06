package main

import "fmt"

func main() {

	// 声明一个长度为10的byte数组，并且赋值
	var arr = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明2个slice，声明方式跟数组一样；区别在于不需要声明长度
	// var slice1 , slice2 []byte
	var slice1 []byte
	var slice2 []byte

	// 赋值给slice1; 注意是冒号
	slice1 = arr[2:5] // slice1=c,d,e (len = 3 , cap = 8)
	fmt.Println("slice1 len", len(slice1))
	fmt.Println("slice1 cap", cap(slice1))
	for key, value := range slice1 {
		fmt.Println("slice1 key", key)
		fmt.Println("slice1 value", value)
	}

	// 赋值给slice2; 注意是冒号
	slice2 = arr[3:4] // slice2=d (len = 1 , cap = 7)
	fmt.Println("slice2 len", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))
	for key, value := range slice2 {
		fmt.Println("slice2 key", key)
		fmt.Println("slice2 value", value)
	}

	//从slice中声明slice
	sliceFromSlice := slice1[1:3] // sliceFromSlice=d,e (len = 2 , cap = 7)
	fmt.Println("sliceFromSlice len", len(sliceFromSlice))
	fmt.Println("sliceFromSlice cap", cap(sliceFromSlice))
	for key, value := range sliceFromSlice {
		fmt.Println("sliceFromSlice key", key)
		fmt.Println("sliceFromSlice value", value)
	}

	// 对slice的slice可以在cap范围内扩展，此时sliceFromSlice2包含
	sliceFromSlice2 := slice1[1:5] // sliceFromSlice2=d,e,f,g (len = 4 , cap = 7)
	fmt.Println("sliceFromSlice2 len", len(sliceFromSlice2))
	fmt.Println("sliceFromSlice2 cap", cap(sliceFromSlice2))
	for key, value := range sliceFromSlice2 {
		fmt.Println("sliceFromSlice2 key", key)
		fmt.Println("sliceFromSlice2 value", value)
	}

	// TODO append 、copy函数
}
