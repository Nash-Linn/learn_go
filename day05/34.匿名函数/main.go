package main

import (
	"fmt"
	"reflect"
)

/*
	匿名函数
	func(参数列表)(返回参数列表){
		函数体
	}
*/

/*
	Go语言不支持在函数内部声明普通函数，只能声明匿名函数
*/

func main() {
	//var f = func() {
	//	fmt.Println("hello func")
	//}
	//f()
	(func(x, y int) {
		fmt.Println(x + y)
	})(10, 14)

	foo()
}

func foo() {
	var bar func(int, int) int

	bar = func(x, y int) int {
		return 100
	}

	fmt.Println(reflect.TypeOf(bar))
}
