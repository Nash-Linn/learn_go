package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	fmt.Printf("s地址%p\n", s)
	foo(s)
	fmt.Println("foo", s)

	var s1 = []int{1, 2, 3}
	fmt.Printf("s1地址%p\n", s)
	foo1(s1)
	fmt.Println("foo1", s1)
}

/*
	因为 s 存的的是 数组 [1,2,3] 的起始地址，长度，容量
	将s[0]改为 100 是在同一个地址里操作
	所以 s 最终值变为 [100,2,3]
*/
func foo(s []int) {
	fmt.Printf("foo s地址 前%p\n", s)
	s[0] = 100
	fmt.Printf("foo s地址 后%p\n", s)
}

/*
	因为本来容量为3 现在增加了一个元素4，切片发生了二倍扩容 生成一个容量变为6的切片 指向了新数组
	函数内 s 的地址发生了变化，操作s[0] = 100 已经是去操作新的数组了
	所以不会影响到外部的s1
	所以 s1 的最终值不受影响为 [1,2,3]
*/
func foo1(s []int) {
	fmt.Printf("foo1 s地址 前%p\n", s)
	s = append(s, 4)
	fmt.Printf("foo1 s地址 后%p\n", s)
	s[0] = 100
}
