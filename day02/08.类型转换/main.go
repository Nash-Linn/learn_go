package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 整型之间的转换

	var x int8 = 100
	var y int16 = 200
	//fmt.Println(x + y)   //报错 无效运算: x + y(类型 int8 和 int16 不匹配)

	fmt.Println(x + int8(y)) // 结果变为44  超出范围 int8 范围[-127,128]

	fmt.Println(int16(x) + y)

	// 字符串与整型之间的转换
	// 字符串转整型
	var ageStr = "32"
	var age, _ = strconv.Atoi(ageStr)
	fmt.Println(age + 1)

	// 整型转字符串
	price := 100
	var priceStr = strconv.Itoa(price)
	fmt.Println(priceStr)

	// strconv Parse 系列函数
	// 将字符串转换成整型   bitSize 限制转换后值的范围
	var ret, _ = strconv.ParseInt("28", 10, 8)
	fmt.Println(ret, reflect.TypeOf(ret))

	// 将字符串转换成浮点型
	var ret1, _ = strconv.ParseFloat("3.1415926", 64)
	fmt.Println(ret1, reflect.TypeOf(ret1))

	// 将字符串转换成布尔值
	var b, _ = strconv.ParseBool("0")
	fmt.Println(b)
	var b1, _ = strconv.ParseBool("-1")
	fmt.Println(b1)
	var b2, _ = strconv.ParseBool("true")
	fmt.Println(b2)
	var b3, _ = strconv.ParseBool("T")
	fmt.Println(b3)
}
