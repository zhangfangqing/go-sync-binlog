package main

import "fmt"

// 声明一个函数类型（类似jdk8的行为参数化）
type testInt func(int) (bool)

//判断奇数函数
func isOdd(a int) (bool) {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

//判断偶数
func isEven(a int) (bool) {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

// 过滤函数请求参数，接受一个函数；（Java8行为参数化）
func filter(slice []int, f testInt) ([]int) {
	// 声明一个slice
	var result []int
	for _, value := range slice {
		if (f(value)) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println(slice)

	odd := filter(slice, isOdd)
	fmt.Println("odd=", odd)

	even := filter(slice, isEven)
	fmt.Println("even=", even)

}
