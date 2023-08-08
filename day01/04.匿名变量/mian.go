package main

import "fmt"

func main() {
	//var a, _ = 1, 2 // _ 为匿名变量
	//fmt.Println(a)

	var a, _ = foo()
	fmt.Println(a)
}

// 当一个函数返回多个值的时候，而我们只需要其中部分值，不要的值可以用匿名变量
func foo() (int, int) {
	return 1, 2
}
