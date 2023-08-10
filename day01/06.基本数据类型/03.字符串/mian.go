package main

import "fmt"

func main() {
	// 字符串是最基本也是最常用的数据类型，是通过 双引号 将多个字符串联起来的一种数据，用于展示文本

	var s string //默认为 ""
	s = "hello world"
	fmt.Println(s)

	// 索引和切片 字符串[索引] 字符串[索引start:end] 左闭右开
	fmt.Println(string(s[1]))

	fmt.Println(s[1:3])
	fmt.Println(s[:7])
	fmt.Println(s[7:])

	// 字符串拼接
	var s1 = "Hi"
	var s2 = " world"
	fmt.Println(s1 + s2)

	// 转义符号
	fmt.Println("Hi\nworld")

	var s3 = "D:\\next\\go.exe"
	fmt.Println(s3) //  D:\next\go.exe

	var s4 = "his name is \"Bob\""
	fmt.Println(s4)

	// 多行打印
	fmt.Println("1.购买血药")
	fmt.Println("2.购买武器")
	fmt.Println("3.生命值槽")

	fmt.Println(`
	1.购买血药
	2.购买武器
	3.生命值槽
    `)
}
