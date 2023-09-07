package main

import "fmt"

func main() {
	// 插入
	var a = []int{1, 2, 3}

	// 在切片 a 最前面插入 0
	a = append([]int{0}, a...)

	fmt.Println(a)
	// 在切片 a 前面插入一个切片
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a)

	// 在切片 b 下标 index 后面插入元素
	var b = []int{1, 2, 3, 4, 5}
	var index = 2

	b = append(b[:index], append([]int{7, 8, 9}, b[index:]...)...)
	fmt.Println(b)

	// 删除元素
	var c = []int{1, 2, 3, 4, 5}
	var cIndex = 2
	c = append(c[:cIndex], c[cIndex+1:]...)
	fmt.Println(c)
}
