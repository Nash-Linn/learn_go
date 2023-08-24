package main

import "fmt"

func main() {
	// 科学运算 + - * / %
	var x, y = 10, 20
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// 关系运算符 > < == != >= <=
	fmt.Println(x > y)

	// 逻辑运算
	// 与运算 && 有假即假
	// 或运算 || 有真即真
	// 非运算 !  取反
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true || false)

	// 赋值运算   += -+ *= /+  ...
	var z = 10
	//z = z + 1 // 自加1
	z += 1 // 相当于  x = x + 1
	z++
	fmt.Println(z)

	// 优先级
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	var t = a + b
	fmt.Println(t)
}
