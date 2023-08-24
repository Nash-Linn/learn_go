package main

import "fmt"

func foo() {
	var y = 100
	fmt.Println(y)
}

func main() {

	var x = 10
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		fmt.Println(x) //内部可以拿到外部的 x
	}
	//fmt.Println(i)   //拿不到 i
}
