package main

import "fmt"

func main() {

	// 循环语句中嵌套分支语句
	for num := 1; num <= 100; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}

	// 分支语句中嵌套循环语句
	var count int
	fmt.Println("请输入一个正整数：")
	fmt.Scan(&count)

	if count > 100 {
		// 从小到大打印 1 - num
		for i := 1; i < count; i++ {
			fmt.Println(i)
		}
	} else {
		// 从大到小打印num - 1
		for i := count; i > 0; i-- {
			fmt.Println(i)
		}
	}
}
