package main

import (
	"fmt"
	"reflect"
)

func main() {
	addToN(10) //实参
	printLing(5)

	add(1, 2, 3, 4)
}

func addToN(n int) { //形参
	var ret = 0
	for i := 1; i <= n; i++ {
		ret += i
	}
	fmt.Println(ret)
}

func printLing(n int) {
	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func add(x ...int) {
	fmt.Println(x, reflect.TypeOf(x))
}
