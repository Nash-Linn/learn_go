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

	fmt.Println(students["name"])
	fmt.Println(students["age"])

	fmt.Println("len", len(students))

	//写入一个 key-value
	students["gender"] = "0"
	fmt.Println(students["gender"])

	//修改一个key-value
	students["name"] = "Bob"
	fmt.Println(students["name"])

	//删除一个 key-value
	delete(students, "gender")
	fmt.Println(students)

	//基于 make 函数声明初始化
	//var students02 map[string]string
	//students02["name"] = "Bob"  // 会报错 还没有空间

	var students02 = make(map[string]string)
	students02["name"] = "Bob"
	fmt.Println(students02)

	// 要想让值有更多类型 需要使用 interface
	var students03 = make(map[string]interface{})
	students03["name"] = "Ken"
	students03["age"] = 30
	fmt.Println(students03)
}
