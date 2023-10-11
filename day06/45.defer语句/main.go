package main

import (
	"fmt"
	"os"
)

/*
	defer语句是go语言提供的一种延迟执行机制，可以在函数退出时执行某个语句或函数。

	defer语句注册了一个函数调用，这个调用会延迟到defer语句所在的函数执行完毕后执行，
	所谓执行完毕是指该函数执行了return语句、函数体已执行完最后一条语句或者发生了panic。


	当执行defer语句时，函数调用不会马上发生，会先把defer注册的函数及变量拷贝到defer栈中保存，知道函数return前才会执行defer中的函数调用。

	**需要额外注意的是**

	这一拷贝拷贝的是那一刻函数的值和参数的值，注册之后再修改函数值或参数时，不会生效
*/

func test01() {
	fmt.Println("test01")
	defer fmt.Println("test02")
	fmt.Println("test03")

	/*
		打印结果：
		test01
		test03
		test02
	*/
}

func test02() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	defer f.Close() // 关闭文件  延迟注册关闭文件操作 保证文件一定关闭
	// 各种文件操作
}

func test03() {
	fmt.Println("test001")
	defer fmt.Println("test002")
	fmt.Println("test003")
	defer fmt.Println("test004")
	fmt.Println("test005")

	/*
			打印结果：
			test001
			test003
			test005
			test004
			test002

			执行顺序：
		    defer 后进先出 先注册的后执行
	*/
}

func test04() {
	foo := func() {
		fmt.Println("foo1")
	}
	defer foo()

	foo = func() {
		fmt.Println("foo2")
	}
}

func test05() {
	x := 10
	defer func(a int) {
		fmt.Println("test05", a)
	}(x)

	x++
}

func test06() {
	x := 10
	defer func() {
		fmt.Println("test06", x)
	}()
	x++
}

func main() {
	test01()
	fmt.Println("-----------------------------------")
	test03()
	fmt.Println("-----------------------------------")
	test04()
	fmt.Println("-----------------------------------")
	test05()
	fmt.Println("-----------------------------------")
	test06()
}
