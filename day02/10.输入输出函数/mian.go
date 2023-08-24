package main

import "fmt"

func main() {
	// 输出函数
	// 1. Print 和 Println
	var name, age = "Nash", 22
	fmt.Println("hello world")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("姓名：", name, "年龄：", age)

	//不换行
	fmt.Print(name)
	fmt.Print(age)
	fmt.Print("\n")

	// Printf : 标准化输出
	var isMarried = false
	fmt.Printf("姓名：%s，年龄：%d ，婚否：%t", name, age, isMarried)

	fmt.Print("\n")
	// Sprintf: 可以返回一个字符串
	var s = fmt.Sprintf("姓名：%s，年龄：%d ，婚否：%t", name, age, isMarried)
	fmt.Println(s)

	fmt.Println("---------------------------------------")

	// 输出函数 IO函数
	// fmt.Scan

	var name1 string
	var age1 int
	fmt.Println("请输入姓名和年龄，用空格分割：")
	scan, err := fmt.Scan(&name1, &age1)
	if err != nil {
		return
	} // 等待用户在命令行输入一个值
	fmt.Println("scan", scan)
	fmt.Println("姓名：", name1)
	fmt.Println("年龄：", age1)
	fmt.Println("end")

	var a, b int
	fmt.Println("请按指定格式输入：")
	fmt.Scanf("%d+%d", &a, &b)
	fmt.Println(a + b)
}
