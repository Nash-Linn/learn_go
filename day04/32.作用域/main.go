package main

import "fmt"

var x = 1

func main() {
	foo()
	fmt.Println(x)
}

func foo() {
	var x int
	x = 10
	fmt.Println(x)

	// if  for 都可以开辟作用域

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
