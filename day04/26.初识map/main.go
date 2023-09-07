package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		var map_name map[key_type]value_type
		map_name 	为 map 的变量名
		key_type 	为键类型
		value_type	为键对应的值类型
	*/
	var stus map[string]string
	fmt.Println(stus, reflect.TypeOf(stus)) //map[] map[string]string

	// 声明并初始化

	var students = map[string]string{
		"name": "Nash",
		"age":  "12",
	}
	fmt.Println(students)

}
