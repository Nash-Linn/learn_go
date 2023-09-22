package main

import (
	"fmt"
	"reflect"
)

/*
	一个高阶函数应该具备下面至少一个特点：
		将一个或多个函数作为形参
		返回一个函数作为其结果
*/

func foo() {
	fmt.Println("foo")
}

// 将函数作为参数传入
func bar(f func()) {
	//经过判断，执行接收到的函数
	f()
}

func main() {
	fmt.Println(foo, reflect.TypeOf(foo))

	bar(foo)
}
