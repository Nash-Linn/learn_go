package main

import "fmt"

func main() {
	test01()
	test02()
}

func test01() {
	fmt.Println("test01-----------------------------------")
	//map 嵌套 map
	stu01 := map[string]string{
		"name": "Bob",
		"age":  "22",
	}
	stu02 := map[string]string{
		"name": "Ken",
		"age":  "23",
	}
	stu03 := map[string]string{
		"name": "Nash",
		"age":  "24",
	}

	var stus = make(map[int]map[string]string)

	stus[1001] = stu01
	stus[1002] = stu02
	stus[1003] = stu03

	fmt.Println(stus)

	//打印 1002 的学生的 年龄
	fmt.Println(stus[1002]["age"])
	fmt.Println("-----------------------------------")
	//循环打印每个学生的 学号 姓名 年龄
	for stuId, stuInfo := range stus {
		fmt.Println(stuId, stuInfo["name"], stuInfo["age"])
	}
	fmt.Println("-----------------------------------")
}

func test02() {
	fmt.Println("test02-----------------------------------")
	//切片嵌套 map
	//stu01 := map[string]string{
	//	"name": "Bob",
	//	"age":  "22",
	//}
	//stu02 := map[string]string{
	//	"name": "Ken",
	//	"age":  "23",
	//}
	//stu03 := map[string]string{
	//	"name": "Nash",
	//	"age":  "24",
	//}

	//var stus = make([]map[string]string, 3)

	//var stus = []map[string]string{stu01, stu02, stu03}

	var stus = []map[string]string{
		{
			"name": "Bob",
			"age":  "22",
		}, {
			"name": "Ken",
			"age":  "23",
		}, {
			"name": "Nash",
			"age":  "24",
		},
	}

	fmt.Println(stus)

	//打印第二个学生的姓名
	fmt.Println(stus[1]["name"])

	//循环打印每个学生的姓名和年龄
	for _, v := range stus {
		fmt.Println(v["name"], v["age"])
	}

}
