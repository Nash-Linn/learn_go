package main

import (
	"fmt"
)

func main() {
	var s []int
	s1 := append(s, 1)
	fmt.Println(s1)

	s2 := append(s1, 2, 3, 4)
	fmt.Println(s2)

	var t = []int{5, 6, 7}
	s3 := append(s2, t...) // t... 能将t内的元素取出来
	fmt.Println(s3)        // [1,2,3,4,5,6,7]

	var s4 = make([]int, 3, 10)
	s5 := append(s4, 100)
	fmt.Println(s5) // [0,0,0,100]
}
