package main

import "fmt"

func main() {
	// 声明一个map; map[keyType]valueType
	nameAgeMap := make(map[string]int)
	//赋值
	nameAgeMap["wangwu"] = 11
	nameAgeMap["liliu"] = 22
	fmt.Println("liliu old age: ", nameAgeMap["liliu"])
	// 修改值
	nameAgeMap["liliu"] = 33
	fmt.Println("liliu new age: ", nameAgeMap["liliu"])

	//初始化一个map
	priceMap := map[string]float32{"java": 4.5, "go": 10, "php": 3}
	// map返回2个值；第一个为value，第二个返回OK；如果key存在返回true
	value1, ok1 := priceMap["go"]
	fmt.Println("val=", value1)
	fmt.Println("ok=", ok1)

	value2, ok2 := priceMap["go2"]
	fmt.Println("val=", value2)
	fmt.Println("ok=", ok2)

	// 删除key
	delete(priceMap, "go")
	value3, ok3 := priceMap["go"]
	fmt.Println("val=", value3)
	fmt.Println("ok=", ok3)

}
