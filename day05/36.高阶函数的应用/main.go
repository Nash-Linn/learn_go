package main

import (
	"fmt"
	"time"
)

func foo() {
	fmt.Println("foo功能开始")
	time.Sleep(time.Second * 2) //模拟程序运行了2s
	fmt.Println("foo功能结束")
}

func bar() {
	fmt.Println("bar功能开始")
	time.Sleep(time.Second * 3) //模拟程序运行了3s
	fmt.Println("bar功能结束")
}

func main() {
	// 补充时间戳

	//f1 := time.Now().Unix()
	//fmt.Println(f1, reflect.TypeOf(f1))
	//
	//time.Sleep(time.Second * 3) //模拟程序运行了3s
	//
	//f2 := time.Now().Unix()
	//fmt.Println(f2, reflect.TypeOf(f2))

	foo()
	bar()

	calcFuncCostTime(foo)
}

// 计算函数运行时间
func calcFuncCostTime(f func()) {
	f1 := time.Now().Unix()
	f()
	f2 := time.Now().Unix()
	costTime := f2 - f1
	fmt.Printf("该函数运行了%d秒\n", costTime)
}
