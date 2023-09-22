package main

import "fmt"

func main() {
	//函数调用
	printRange(1, 10) //按照位置传参的实际参数，称为实参

	sum := add(2, 3)
	fmt.Println(sum)

	foo()
}

// 函数声明 函数名尽量 见名知意
func printRange(start int, end int) { //start ,end 称为形参
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func add(x, y int) int {
	return x + y //返回值
}

//返回值命名
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return //不写返回参数 会默认返回 sum ,sub
}

//作用域
var x = 100 //全局变量
func foo() {
	fmt.Println(x) //可以在函数内获取到全局变量

	var foo1 = func() {
		var y = 10
		fmt.Println(y)
	}
	foo1()

	//fmt.Println(y) //获取不到 foo1 内部的变量y

}
