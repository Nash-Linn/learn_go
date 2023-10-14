package main

import "fmt"

type Addr struct {
	province string
	city     string
	country  string
}

type Student struct {
	string //匿名字段  相当于 string  string
	name   string
	age    int
	Addr   //匿名字段
}

func main() {
	s := Student{"string", "Nash", 18, Addr{"广东", "深圳", "南山区"}}
	fmt.Println(s)

	fmt.Println(s.Addr.province)
}
