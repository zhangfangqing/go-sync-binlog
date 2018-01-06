package main

import "fmt"

const (
	a = iota
	b
	c
	d
)

func main() {
	const e = iota
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
