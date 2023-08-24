package main

import "fmt"

func main() {
	// 1. 先声明再赋值
	// 数组的声明
	// 数组必须限制长度
	var arr [3]int
	fmt.Println(arr)

	// 索引赋值
	arr[0] = 10
	fmt.Println(arr)

	// 2.数组的声明并赋值
	var names = [3]string{"Nash", "Bob", "Mike"}
	fmt.Println(names) // [Nash Bob Mike]
	var ages = [3]int{22, 23, 24}
	fmt.Println(ages)

	// 3.省略长度赋值  使用 ... 代替
	var fruits = [...]string{"apple", "peach", "pear"}
	fmt.Println(fruits)
}
