package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test01()
	test02()
	test03()
	//test04()

	//test05()

	test06()
}

//判断一个整型数字是奇数还是偶数
func test01() {
	var x = 10
	fmt.Println(x%2 == 0)
}

//下面代码执行会报什么错误，如何正确计算a+b的值
func test02() {
	var a int8 = 100
	var b int16 = 200
	//fmt.Println(a + b) // 变量类型不同
	fmt.Println(int16(a) + b)
}

// 下面代码的执行结果是
func test03() {
	cmd := ""
	cmd += "l"
	cmd += "s"
	fmt.Println(cmd)
}

//引导用户输入生日字符串，格式为“年-月-日”，比如“1990-3-16”，然后按
//“您的生日是1990年-3月-16日”的格式化字符串输入到终端控制台
func test04() {
	fmt.Println("请按照格式为“年-月-日”，比如“1990-3-16”，输入您的生日：")
	var birth string //birth ""
	fmt.Scan(&birth)
	s := strings.Split(birth, "-")
	fmt.Printf("您的生日是%s年-%s月-%s日", s[0], s[1], s[2])
}

//引导用户输入一个名字，判断该名字是否以小写a或者大写A开头，如果是打印true，否则打印false
func test05() {
	var name string
	fmt.Print("英文名：")
	fmt.Scan(&name)
	//直接判断
	//fmt.Println(strings.HasPrefix(name, "a") || strings.HasPrefix(name, "A"))

	//先进行转换
	upperName := strings.ToUpper(name)
	fmt.Println(strings.HasPrefix(upperName, "A"))
}

// 将数字100转为字符串100
func test06() {
	var x = 100
	fmt.Println(strconv.Itoa(x))
}
