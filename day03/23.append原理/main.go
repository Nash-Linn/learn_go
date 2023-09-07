package main

import "fmt"

func main() {
	a := []int{11, 22, 33}
	fmt.Println(len(a), cap(a)) // 3, 3

	c := append(a, 44) // 由于 切片 a 容量是3且长度是3已经满了，进行append会开辟一个新空间，并对容量进行两倍扩容（将切片a容量扩大到原来的两倍）
	a[0] = 100
	fmt.Println(a)              // [100,22,33]
	fmt.Println(len(a), cap(a)) //3,3
	fmt.Println(c)              // [11,22,33,44]
	fmt.Println(len(c), cap(c)) //4,6

	// 容量够的情况
	p := make([]int, 3, 10)
	fmt.Println(p) // [0,0,0]
	q := append(p, 11, 12)
	fmt.Println(p) //[0,0,0]
	fmt.Println(q) //[0,0,0,11,12]
	p[0] = 100
	fmt.Println(p) //[100,0,0]
	fmt.Println(q) //[100,0,0,11,12]

	fmt.Println("-------------------------------------------")
	//
	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2] // [10,20]
	s2 := s1       // [10,20]
	s3 := append(append(append(s1, 1), 2), 3)
	//  append(append([10,20,1], 2), 3)
	// 	append([10,20,1,2], 3)
	//  [10,20,1,2,3]
	s1[0] = 1000
	fmt.Println(s1)  //[1000,20]
	fmt.Println(s2)  //[1000,20]
	fmt.Println(s3)  //[10,20,1,2,3]
	fmt.Println(arr) //[1000,20,1,2]
}
