package main

import "fmt"

func main() {
	//var s = []int{1, 2, 3}
	//s[0] = 100 // 已经对切片 s 声明并初始化 所以可修改 s[0]

	//var s []int
	//s[0] = 100 //不可赋值 前面只是对 s 声明 未开辟底层数组

	//var s = make([]int, 5, 10)
	//s[0] = 100 //可修改

	var emps = make([]string, 3, 5)
	emps[0] = "张三"
	emps[1] = "李四"
	emps[2] = "王五"
	fmt.Println(emps) //[张三 李四 王五]
	emps2 := append(emps, "rain")
	fmt.Println(emps2) //[张三 李四 王五 rain]
	emps3 := append(emps2, "eric")
	fmt.Println(emps3) //[张三 李四 王五 rain eric]
	fmt.Println(emps2) //[张三 李四 王五 rain]

	//容量不够时发生二倍扩容
	emps4 := append(emps3, "nash")
	fmt.Println(emps4)                  //[张三 李四 王五 rain eric nash]
	fmt.Println(len(emps4), cap(emps4)) // 6 10

	fmt.Println("-------------------------------------")

	var s = [3]int{1, 2, 3}
	s1 := s[:]
	s2 := append(s1, 4)
	fmt.Println(s)  //[1 2 3]
	fmt.Println(s1) //[1 2 3]
	fmt.Println(s2) //[1 2 3 4]
}
