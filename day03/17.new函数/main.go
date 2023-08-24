package main

import "fmt"

func main() {
	// 基本数据类型 （整型，浮点型，字符串，布尔类型）范围属于值类型
	// 值类型特点，声明未赋值之前存在一个默认值（zero value）
	var x int
	fmt.Println(x)

	var name string
	fmt.Println(name)

	// 指针类型属于引用类型，包括切片，map，channel都是引用类型
	// 引用类型当声明未赋值之前是没有开辟空间的，即没有默认值
	var p *int
	p = new(int) // 用new 函数开辟一个空间

	*p = 10
	fmt.Println(p)
}
