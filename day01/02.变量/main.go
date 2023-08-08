package main

import "fmt"

func main() {
	// 声明变量  var 变量名 类型

	var age int //声明未赋值，默认 零 值
	fmt.Println(age)
	age = 22
	fmt.Println(age)
	age = 30
	fmt.Println(age)

	// 声明并赋值一行实现
	var name string = "张三"
	fmt.Println(name)

	// 声明并赋值的简洁写法  只能在函数内部使用 全局变量不能用这种写法
	name2 := "李四"
	fmt.Println(name2)

	// 一行声明多个变量
	var x, y int
	fmt.Println(x, y)

	var (
		a int    // 默认值 0
		b string // 默认值 空字符串
		c bool   // 默认值 false
	)
	fmt.Println(a, b, c)

	// 一行声明赋值多个变量

	//var o, p, q = "o", "p", "q"
	o, p, q := "o", "p", "q"

	fmt.Println(o, p, q)
}
