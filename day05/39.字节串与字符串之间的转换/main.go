package main

import "fmt"

func main() {
	var msg = "嗨abc"

	// 字符串转换成字节串
	fmt.Println([]byte(msg)) //[229 151 168 97 98 99]

	fmt.Println([]rune(msg)) //[21992 97 98 99]

	// 字节串转换成字符串
	info1 := []byte(msg)
	info2 := []byte{229, 151, 168, 97, 98, 99}

	fmt.Println("info1", info1)
	fmt.Println("info2", info2)
	fmt.Println("info1Str", string(info1))
	fmt.Println("info2Str", string(info2))
}
