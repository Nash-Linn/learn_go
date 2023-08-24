package main

import (
	"fmt"
	"reflect"
)

func main() {
	// &变量 ： 获取变量地址
	var x = 10
	fmt.Printf("赋值之前x的对应地址%p\n", &x)
	x = 100
	fmt.Printf("赋值之后x的对应地址%p\n", &x)

	var y string
	y = "hello"
	fmt.Println("y", y)

	var p *int // p 是一个整型指针类型
	p = &x     //  var p = &x
	fmt.Println("p", p)

	// *指针变量：通过取值操作 获取地址存的值
	fmt.Println("p这个地址里的值:", *p, ",类型：", reflect.TypeOf(*p))
	*p = 1000
	fmt.Println(x)

	/*
		var a = 1
		var b = a
		b = 100 // 改了b 不会影响 a
		fmt.Println(a, b)
	*/

	var a = 1
	var b = &a
	var c *int
	c = b

	*b = 100 //  a  *b  *c 值相同
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)

	// 案例1
	var a1 = 100
	var b1 = &a1    // a1的地址
	var c1 = &b1    // b1的地址
	**c1 = 200      // *c1 拿到 b1的地址里存的值 即 a1 的地址 **c1拿到 a1地址里的值 然后取修改为 200
	fmt.Println(a1) // 200

	// 案例2
	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)  //2
	fmt.Println(*p2) //2
}
