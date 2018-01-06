package main

import "fmt"

type student struct {
	name string
	age  int
}

//定义函数返回年龄大的人、年龄差
func getMaxStudent(stu1, stu2 student) (student, int) {
	if stu1.age > stu2.age {
		return stu1, stu1.age - stu2.age
	} else {
		return stu2, stu2.age - stu1.age
	}
}

func main() {
	var s1 student
	s1.name = "fangqing"
	s1.age = 29

	s2 := student{"fangshu", 26}

	s3 := student{name: "shumei", age: 24}

	s12, age12 := getMaxStudent(s1, s2)
	fmt.Println("s1 s2 compare , student=", s12)
	fmt.Println("s1 s2 compare, age=", age12)

	s13, age13 := getMaxStudent(s1, s3)
	fmt.Println("s1 s3 compare , student=", s13)
	fmt.Println("s1 s3 compare, age=", age13)

	s23, age23 := getMaxStudent(s2, s3)
	fmt.Println("s2 s3 compare , student=", s23)
	fmt.Println("s2 s3 compare, age=", age23)
}
