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
}
