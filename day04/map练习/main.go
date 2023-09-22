package main

import (
	"fmt"
	"strconv"
)

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

	//添加一个学生
	newMap := map[string]string{"name": "Eric", "age": "30"}
	stus = append(stus, newMap)
	fmt.Println(stus)

	//删除一个学生
	//a = append(a[:index], a[index + 1:]...)

	//stus = append(stus[:1], stus[2:]...)
	//fmt.Println("stus", stus)

	// 删除学生eric的map
	// 查询eric的索引位置
	var deleteIndex = 0
	for index, stuMap := range stus {
		if stuMap["name"] == "Eric" {
			deleteIndex = index
		}
	}

	stus = append(stus[:deleteIndex], stus[1+deleteIndex:]...)
	fmt.Println("stus", stus)

	// 将姓名为Ken的学生的年龄自加一岁

	for _, stuMap := range stus {
		if stuMap["name"] == "Ken" {
			// 类型转换
			age, _ := strconv.Atoi(stuMap["age"])
			stuMap["age"] = strconv.Itoa(age + 1)
		}
	}
}
