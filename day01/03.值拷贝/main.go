package main

import "fmt"

func main() {
	var x = 10
	var y = x //值拷贝  x,y对应的是完全不同的地址空间
	x = 20
	fmt.Println(x, y)

	var a = 1 + 1
	fmt.Println(a)

	var b = x * y
	fmt.Println(b)
}
