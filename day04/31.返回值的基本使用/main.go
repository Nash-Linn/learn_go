package main

import "fmt"

func main() {
	var sum = addToN(5)
	fmt.Println(sum)

	b, user := login("root", "112")
	if b {
		fmt.Println(user)
	}
}

func addToN(n int) int { //形参
	var ret = 0
	for i := 1; i <= n; i++ {
		ret += i
	}
	return ret // 函数的终止语句
}

func login(user, pwd string) (isSuccess bool, userName string) { //多个返回值 也可以给返回值命名
	if user == "root" && pwd == "123" {
		//登录成功
		return true, user
	} else {
		return false, ""
	}
}

func login1(user, pwd string) (isSuccess bool, userName string) { //多个返回值 也可以给返回值命名
	if user == "root" && pwd == "123" {
		//登录成功
		isSuccess = true
		userName = user
	} else {
		isSuccess = false
		userName = ""
	}
	return
}
