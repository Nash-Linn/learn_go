package main

import "fmt"

func main() {
	/*
	   for 表达式 {
	      //循环语句
	   }
	*/
	var count = 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	for count > 0 {
		fmt.Println(count)
		count--
	}

	// 三要素for循环
	/*
	   var count = 0  //初始语句
	   for count < 10 { // 条件表达式
	       fmt.Println(count)
	       count++ //步进语句
	   }
	*/

	/*
	   for num := 0; num < 10; num++ {
	     // 循环语句
	     fmt.Println(num)
	   }
	*/

	// 基于循环  1+2+3+4+...+100 ?

	var total = 0
	for num := 1; num <= 100; num++ {
		total += num
	}
	fmt.Println("total", total)
}
