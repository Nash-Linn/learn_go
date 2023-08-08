package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1.整型  int

	/*
	 int8：一个字节 范围 [-127 ~ 128]   2的8次方
	 ...
	*/
	var x int8 = 100
	fmt.Println(x)
	//x = 200 超出范围报错

	// 2.浮点型
	var f1 float32 = 3.012345678912345678
	var f2 float64 = 3.012345678912345678
	fmt.Println(f1)                 //3.0123458
	fmt.Println(reflect.TypeOf(f1)) //float32
	fmt.Println(f2)                 //3.0123456789123457
	fmt.Println(reflect.TypeOf(f2)) //float64

}
