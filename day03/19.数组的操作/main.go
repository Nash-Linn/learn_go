package main

import (
	"fmt"
	"reflect"
)

func main() {
	var names = [3]string{"Nash", "Bob", "Mike"}

	// 1.索引操作 数组[索引]
	fmt.Println(names[2])
	names[2] = "Alvin"
	fmt.Println(names)

	// 2.切片操作 数组[start索引:end索引] 范围左闭右开
	var arr = [5]int{11, 12, 13, 14, 15}
	s := arr[1:3]
	fmt.Println(s, reflect.TypeOf(s)) //[12 13] []int

	// 3.遍历数组
	var arr2 = [5]int{20, 21, 22, 23, 24}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	//range循环
	for i, v := range arr2 {
		fmt.Println(i)
		fmt.Println(v)
	}
}
