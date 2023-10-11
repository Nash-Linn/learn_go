package main

import (
	"fmt"
	"reflect"
)

func main() {

	// uint8 只能存0-255
	var x uint8
	x = 255
	fmt.Println(x)

	//字符
	y := 'a' // 双引号标识字符串，单引号标识字符
	fmt.Println("y", y, reflect.TypeOf(y))

	// byte
	// type byte = int8
	var b byte // byte字节类型：声明字符
	b = 'a'
	fmt.Println("b", b, reflect.TypeOf(b))

	//字符串
	/*
			go语言的string是一种数据类型，这个数据类型占用16字节空间，
			前8字节是一个指针，指向字符串值的地址，
		    后8字节是一个整数，标识字符串的长度
	*/

	// rune类型
	// type rune = int32
	var z rune
	z = '严'
	fmt.Printf("字符’严‘unicode的十进制：%d\n", z)  //20005
	fmt.Printf("字符’严‘unicode的十六进制：%x\n", z) //4e25
	fmt.Printf("字符’严‘unicode的二进制：%b\n", z)  //100111000100101

	// unicode utf8转换
	/*
		| Unicode 符号范围(16进制)| UTF-8编码方式(2进制)
		| -----------------------|--------------------------------------------------------------
		| 0000 0000  - 0000 007F |                                                    0xxx xxxx
		| 0000 0000  - 0000 07FF |                                          110x xxxx 10xx xxxx
		| 0000 0000  - 0000 FFFF |                                1110 xxxx 10xx xxxx 10xx xxxx
		| 0001 0000  - 0010 FFFF |                      1111 0xxx 10xx xxxx 10xx xxxx 10xx xxxx
		| 0020 0000  - 03FF FFFF |            1111 10xx 10xx xxxx 10xx xxxx 10xx xxxx 10xx xxxx
		| 0400 0000  - 7FFF FFFF |  1111 110x 10xx xxxx 10xx xxxx 10xx xxxx 10xx xxxx 10xx xxxx
	*/
	/*
		例子：以汉字"严"为例, 演示如何实现UTF-8编码 - 已知"严"的unicode是4E25(1001110 00100101),
		根据上表, 可以发现4E25处在第三行的范围内(0000 0800 - 0000 FFFF),
		因此"严"的UTF-8编码需要三个字节, 即格式是 "1110xxxx 10xxxxxx 10xxxxxx".
		然后, 从"严"的最后一个二进制位开始, 依次从后向前填入格式中的x, 多出的位补0.
		这样就得到了, "严"的UTF-8编码是 "11100100 10111000 10100101", 转换成十六进制就是E4B8A5.
	*/

	var s = "严abc"
	fmt.Println(s, len(s)) //len 字节

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	/*
		228
		184
		165
		97
		98
		99
		前三个编码代表’严‘，后面三个代表abc
	*/

	for i, v := range s {
		fmt.Println(i, v)
	}
	/*
		0 20005
		3 97
		4 98
		5 99
	*/
	// 通过长度遍历 按照字节遍历，按照range遍历 按照字符遍历
}
