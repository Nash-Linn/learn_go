package main

import "fmt"

//func getCounter() func() int {
//	var i = 0
//	return func() int {
//		i++
//		fmt.Println(i)
//		return i
//	}
//}

func getCounter(i int) func() int {
	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}

func main() {
	//count := getCounter()

	count := getCounter(1)
	count()
	count()
	c := count()
	fmt.Println("c", c)

	count1 := getCounter(10)
	count1()
	count1()
	count1()
}
