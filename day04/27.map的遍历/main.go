package main

import "fmt"

func main() {
	// 遍历 map 对象
	var s = map[string]string{
		"name": "Bob",
		"age":  "20",
	}

	for k, v := range s {
		fmt.Println(k)
		fmt.Println(v)
	}

	var noSortMap = map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	}

	for k, v := range noSortMap {
		fmt.Println(k, v) //打印出来的顺序是无序的
	}
}
