package main

import "fmt"

func main() {
	// 分支语句

	// 1. 单分支语句
	var name = "root1"

	if name == "root" {
		// if 为真执行的代码块
		fmt.Println("姓名匹配成功")
	}

	// 2. 双分支语句
	/*	var age int
		fmt.Println("请输入您的年龄：")
		fmt.Scan(&age)
		if age >= 18 {
			// 表达式为真执行语句
			fmt.Println("恭喜")
		} else {
			// 表达式为假执行语句
			fmt.Println("很遗憾")
		}
	*/
	// 3. 多分支语句
	var score int
	fmt.Println("请输入你的成绩：")
	fmt.Scan(&score)

	if score < 0 || score > 100 {
		fmt.Println("输入的成绩不合法")
	} else if score > 90 {
		fmt.Println("成绩优秀")
	} else if score > 60 {
		fmt.Println("成绩良好")
	} else {
		fmt.Println("不及格")
	}

	// switch 语句进行分支判断，switch值判断
	var week int
	fmt.Println("请输入星期的数字：")
	fmt.Scan(&week)

	switch week {
	case 0:
		fmt.Println("星期日")
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	}
}
