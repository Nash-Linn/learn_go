package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 构建切片方式一：通过数组切片操作获取切片对象
	var arr = [3]string{"rain", "eric", "alvin"}
	fmt.Println(arr, reflect.TypeOf(arr)) //[rain eric alvin] [3]string

	s1 := arr[0:2]
	fmt.Println(s1, reflect.TypeOf(s1)) //[rain eric] []string

	s2 := arr[1:]
	fmt.Println(s2, reflect.TypeOf(s2))

	s2[0] = "nash"
	fmt.Println(s1)
	fmt.Println(arr)

	// 切片是对数组的引用
	var a = [5]int{1, 2, 3, 4, 5}
	var slice = a[:] // slice 空间里存 起始地址 长度 容量
	fmt.Println(slice)
	newSlice := slice[1:3]
	fmt.Println(newSlice) //  [2,3]

	// 构建切片方式二：直接声明切片
	// 长度 len：切片后元素个数
	// 容量 cap：第一个元素开始数，到其底层数组元素末尾的个数
	var q = []int{10, 11, 12, 13, 14}
	fmt.Println(q)

	q1 := q[1:4]
	fmt.Println(len(q1), cap(q1))

	q2 := q[3:]
	fmt.Println(len(q2), cap(q2))

	q3 := q1[1:3]
	fmt.Println(len(q3), cap(q3))

	// 练习
	test()
}

func test() {
	// 练习1
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1) // [1,2,4]

	// 练习2
	var a = []int{1, 2, 3} // a 空间里存 起始地址 长度 容量
	b := a                 // 将 a 的值赋给 b 所以 b 里也存切片的 起始地址 长度 容量
	a[0] = 100
	fmt.Println(b) //[100,2,3]
}
