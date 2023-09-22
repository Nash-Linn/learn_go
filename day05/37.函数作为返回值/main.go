package main

import "fmt"

func foo() func() int {
	// 声明匿名函数
	var inner = func() int {
		fmt.Println("一个新的函数")
		return 100
	}
	return inner
}

func main() {
	var f func() int
	f = foo() //返回inner函数体赋值给f变量
	f()       // 函数调用
}
