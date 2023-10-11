package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func foo() {
	time.Sleep(time.Second * 2)
	fmt.Println("foo功能")
}

//不影响函数已有功能 添加计数功能
func getCalledNum(count int, f func()) func() {
	return func() {
		f()
		count++
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("%s函数被调用第%d次\n", funcName, count)
	}
}

//不影响函数已有功能 添加计时功能
func getCalledTimer(f func()) func() {
	return func() {
		var timer int64
		start := time.Now().Unix()
		f()
		end := time.Now().Unix()
		timer = end - start
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("%s函数运行了%d秒\n", funcName, timer)
	}
}

func main() {
	calledFoo := getCalledNum(0, foo)
	calledFoo()
	calledFoo()

	calledFooTimer := getCalledTimer(foo)
	calledFooTimer()
}
