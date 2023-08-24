package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		if i == 3 {
			continue //退出当次循环
		}

		if i == 6 {
			break //退出循环
		}

		fmt.Println(i)
	}
}
