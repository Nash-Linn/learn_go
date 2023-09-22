package main

import (
	"fmt"
	"reflect"
)

func main() {

	var info = make(map[string]string)
	info["name"] = "Nash"
	info["age"] = "18"
	fmt.Println(info)
	info2 := info
	info["gender"] = "male"
	fmt.Println(info2)

	fmt.Println("--------------------------------------------------")

	//map增删改查
	test01()

	// 切片 嵌套 map
	test02()

	// 补充数组类型
	test03()

	// map 嵌套 map
	test04()

	// map 嵌套 切片
	test05()
}

func test01() {
	//map增删改查
	var info = make(map[string]string)
	info["name"] = "Nash"
	info["age"] = "18"

	value, exist := info["gender"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("性别不存在")
	}

	info["gender"] = "female" //添加键值对
	info["gender"] = "male"   //修改键值对

	delete(info, "gender")

	for k, v := range info {
		fmt.Println(k, v)
	}
	fmt.Println("--------------------------------------------------")
}

func test02() {
	// 切片 嵌套 map
	var y = []map[string]string{
		{
			"name": "Nash",
			"age":  "33",
		}, {
			"name": "Ken",
			"age":  "22",
		},
	}

	fmt.Println(y)

	fmt.Println("--------------------------------------------------")
}

func test03() {
	// 补充数组类型
	var x = [3]int{1, 2, 3}
	var y = [3]string{"Nash", "Rain", "Eric"}
	fmt.Println(reflect.TypeOf(x)) //[3]int
	fmt.Println(reflect.TypeOf(y)) //[3]string
	// x 和 y 的类型不同
}

func test04() {
	var data = make(map[string]map[string]string)
	data["Nash"] = map[string]string{
		"age":  "22",
		"name": "Nash",
	}
	fmt.Println(data)
	fmt.Println(data["Nash"]["name"])
}

func test05() {
	var data = make(map[string][]string)
	data["山东省"] = []string{"威海", "日照"}
}
