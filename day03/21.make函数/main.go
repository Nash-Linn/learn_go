package main

import "fmt"

func main() {

	//var s []int
	//s[0] = 1 // 直接修改会报错 因为切片 s 未初始化

	//初始化创建空间
	var s = make([]int, 5, 10)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	s[0] = 100
	fmt.Println(s)

	// 练习
	test()
}

func test() {
	a := make([]int, 5)
	b := a[0:3]
	a[0] = 100
	fmt.Println(a) // [100,0,0,0,0]
	fmt.Println(b) // [100,0,0]
}
