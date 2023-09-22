package main

import "fmt"

func main() {
	var x = 1
	foo(x)

	var age = 22
	setAge(age)             //函数内部只是对传入的值进行值拷贝
	fmt.Println("age", age) // 无法影响到外部的age

	setAgeByAddress(&age) //传入age的地址
	fmt.Println("ageA", age)
}

func foo(x int) {
	x = 100
	fmt.Println("foo", x)
}

func setAge(age int) {
	age++
}

//采用地址方式修改
func setAgeByAddress(ageAddress *int) {
	*ageAddress++
}
