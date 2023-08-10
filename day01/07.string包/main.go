package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Nash"
	var newName = strings.ToUpper(name)
	fmt.Println(name)
	fmt.Println(newName)
	fmt.Println(strings.ToLower(newName))

	var s = "rain rain"
	//  strings.HasPrefix() 判断以什么开头
	fmt.Println(strings.HasPrefix(s, "ra"))

	//  strings.HasSuffix() 判断以什么结尾
	fmt.Println(strings.HasSuffix(s, "in"))

	//  strings.Contains() 判断是否包含
	fmt.Println(strings.Contains(s, "iin"))

	// strings.Trim() 去除两端对应的字符
	var username = "   Nash   "
	fmt.Println(strings.Trim(username, " "))

	// strings.TrimLeft() 去除左边对应的字符
	fmt.Println(strings.TrimLeft(username, " "))

	// strings.TrimLeft() 去除右边对应的字符
	fmt.Println(strings.TrimRight(username, " "))

	// strings.TrimSpace() 去除两端空格
	fmt.Println(strings.TrimSpace(username))

	// index:索引
	var s2 = "hello world"
	fmt.Println(strings.Index(s2, "wo"))

	// 分割 拼接
	var s3 = "hello world hi Nash"
	var nameSlice = strings.Split(s3, " ")
	fmt.Println(nameSlice)
	fmt.Println(strings.Join(nameSlice, ""))
}
