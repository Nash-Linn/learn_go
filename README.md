# Golang

## 一.开始

### 1 Go环境安装

#### 1.1 Go编译器下载

官网 https://golang.google.cn/

#### 1.2 配置环境变量 

![image-20230802163400201](README.assets/image-20230802163400201.png)

在该文件夹下 创建 三个文件夹

![image-20230802163454451](README.assets/image-20230802163454451.png)

#### 1.3 hello word

在任意位置，创建一个文件  hello/main.go

```go
package main

import "fmt"

func main()  {
	fmt.Println("hello word")
}
```

然后对文件进行编译 

在当前文件夹打开cmd

```
方法1：
go build mian.go
方法2：
go build
方法3：
go build -o 别名
```

上述方法会生成一个可执行文件

执行 `go run mian.go`  可以编译并执行



#### 1.4 goland 安装

官网 [Download GoLand: A Go IDE with extended support for JavaScript, TypeScript, and databases (jetbrains.com)](https://www.jetbrains.com/go/download/#section=windows)



#### 1.5 创建项目

在之前配置的环境变量文件夹中的 src 下创建项目 learn

![image-20230804162443382](README.assets/image-20230804162443382.png)

新建文件 day01/01helloworld/main.go

![image-20230804162809379](README.assets/image-20230804162809379.png)

```
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
}
```

在01helloworld文件夹下 打开 编译器终端执行  go build

正常执行则会生成 一个后缀为.exe 的文件

如果报错 检查 系统环境变量和用户环境变量 中的 GOPATH

还要修改编译器的设置

![image-20230804163540248](README.assets/image-20230804163540248.png)

还需要关闭  go mod 模式

```
go env -w GO111MODULE=off
```



想要执行还可以直接使用编辑器

![image-20230804163839123](README.assets/image-20230804163839123.png)



## 二.基础语法

### 1.注释

```go
//单行注释

/*
	多
	行
	注
	释
*/
```

### 2.变量

```
var 变量名 变量类型

var x int
var s string
var b bool
```

```go
package main

import "fmt"

func main() {
	// 声明变量  var 变量名 类型

	var age int //声明未赋值，默认 零 值
	fmt.Println(age)
	age = 22
	fmt.Println(age)
	age = 30
	fmt.Println(age)

	// 声明并赋值一行实现
	var name string = "张三"
	fmt.Println(name)

	// 声明并赋值的简洁写法  只能在函数内部使用 全局变量不能用这种写法
	name2 := "李四"
	fmt.Println(name2)

	// 一行声明多个变量
	var x, y int
	fmt.Println(x, y)

	var (
		a int    // 默认值 0
		b string // 默认值 空字符串
		c bool   // 默认值 false
	)
	fmt.Println(a, b, c)

	// 一行声明赋值多个变量

	//var o, p, q = "o", "p", "q"
	o, p, q := "o", "p", "q"

	fmt.Println(o, p, q)
}
```

### 3.值拷贝

```go
package main

import "fmt"

func main() {
	var x = 10
	var y = x  //值拷贝  x,y对应的是完全不同的地址空间
	x = 20
	fmt.Println(x, y)

	var a = 1 + 1
	fmt.Println(a)

	var b = x * y
	fmt.Println(b)
}
```

### 4.匿名变量

```go
package main

import "fmt"

func main() {
	//var a, _ = 1, 2 // _ 为匿名变量
	//fmt.Println(a)

	var a, _ = foo()
	fmt.Println(a)
}

// 当一个函数返回多个值的时候，而我们只需要其中部分值，不要的值可以用匿名变量
func foo() (int, int) {
	return 1, 2
}
```

### 5.语句分隔符

```go
package main

import "fmt"

func main() {
	// ; 和 换行 作为语句分隔符
	// 推荐换行作为分隔符
	var x = 100
	var y = 200
	fmt.Println(x, y)
}
```

### 6.基本数据类型

#### 6.1 整型和浮点型

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// 1.整型  int

	/*
	 int8：一个字节 范围 [-127 ~ 128]   2的8次方
	 ...
	*/
	var x int8 = 100
	fmt.Println(x)
	//x = 200 超出范围报错

	// 2.浮点型
	var f1 float32 = 3.012345678912345678
	var f2 float64 = 3.012345678912345678
	fmt.Println(f1)                 //3.0123458
	fmt.Println(reflect.TypeOf(f1)) //float32
	fmt.Println(f2)                 //3.0123456789123457
	fmt.Println(reflect.TypeOf(f2)) //float64

}
```

#### 6.2 布尔类型

```go
package main

import "fmt"

func main() {
	//布尔类型
	var b bool
	b = true
	b = false
	fmt.Println(b)

	fmt.Println(2 > 1)
}
```

#### 6.3 字符串类型

```go
package main

import "fmt"

func main() {
	// 字符串是最基本也是最常用的数据类型，是通过 双引号 将多个字符串联起来的一种数据，用于展示文本

	var s string //默认为 ""
	s = "hello world"
	fmt.Println(s)

	// 索引和切片 字符串[索引] 字符串[索引start:end] 左闭右开
	fmt.Println(string(s[1]))

	fmt.Println(s[1:3])
	fmt.Println(s[:7])
	fmt.Println(s[7:])

	// 字符串拼接
	var s1 = "Hi"
	var s2 = " world"
	fmt.Println(s1 + s2)

	// 转义符号
	fmt.Println("Hi\nworld")

	var s3 = "D:\\next\\go.exe"
	fmt.Println(s3) //  D:\next\go.exe

	var s4 = "his name is \"Bob\""
	fmt.Println(s4)

	// 多行打印
	fmt.Println("1.购买血药")
	fmt.Println("2.购买武器")
	fmt.Println("3.生命值槽")

	fmt.Println(`
	1.购买血药
	2.购买武器
	3.生命值槽
    `)
}
```

### 7.string包

```go
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
```

### 8.类型转换

```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 整型之间的转换

	var x int8 = 100
	var y int16 = 200
	//fmt.Println(x + y)   //报错 无效运算: x + y(类型 int8 和 int16 不匹配)

	fmt.Println(x + int8(y)) // 结果变为44  超出范围 int8 范围[-127,128]

	fmt.Println(int16(x) + y)

	// 字符串与整型之间的转换
	// 字符串转整型
	var ageStr = "32"
	var age, _ = strconv.Atoi(ageStr)
	fmt.Println(age + 1)

	// 整型转字符串
	price := 100
	var priceStr = strconv.Itoa(price)
	fmt.Println(priceStr)

	// strconv Parse 系列函数
	// 将字符串转换成整型   bitSize 限制转换后值的范围
	var ret, _ = strconv.ParseInt("28", 10, 8)
	fmt.Println(ret, reflect.TypeOf(ret))

	// 将字符串转换成浮点型
	var ret1, _ = strconv.ParseFloat("3.1415926", 64)
	fmt.Println(ret1, reflect.TypeOf(ret1))

	// 将字符串转换成布尔值
	var b, _ = strconv.ParseBool("0")
	fmt.Println(b)
	var b1, _ = strconv.ParseBool("-1")
	fmt.Println(b1)
	var b2, _ = strconv.ParseBool("true")
	fmt.Println(b2)
	var b3, _ = strconv.ParseBool("T")
	fmt.Println(b3)
}
```



### 9.运算符

```
package main

import "fmt"

func main() {
	// 科学运算 + - * / %
	var x, y = 10, 20
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// 关系运算符 > < == != >= <=
	fmt.Println(x > y)

	// 逻辑运算
	// 与运算 && 有假即假
	// 或运算 || 有真即真
	// 非运算 !  取反
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true || false)

	// 赋值运算   += -+ *= /+  ...
	var z = 10
	//z = z + 1 // 自加1
	z += 1 // 相当于  x = x + 1
	z++
	fmt.Println(z)

	// 优先级
	var a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
	var t = a + b
	fmt.Println(t)
}

```

#### 优先级

![image-20230810154214191](README.assets/image-20230810154214191.png)



### 10.输入输出函数

```go
package main

import "fmt"

func main() {
	// 输出函数
	// 1. Print 和 Println
	var name, age = "Nash", 22
	fmt.Println("hello world")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println("姓名：", name, "年龄：", age)

	//不换行
	fmt.Print(name)
	fmt.Print(age)
	fmt.Print("\n")

	// Printf : 标准化输出
	var isMarried = false
	fmt.Printf("姓名：%s，年龄：%d ，婚否：%t", name, age, isMarried)

	fmt.Print("\n")
	// Sprintf: 可以返回一个字符串
	var s = fmt.Sprintf("姓名：%s，年龄：%d ，婚否：%t", name, age, isMarried)
	fmt.Println(s)

	fmt.Println("---------------------------------------")

	// 输出函数 IO函数
	// fmt.Scan

	var name1 string
	var age1 int
	fmt.Println("请输入姓名和年龄，用空格分割：")
	scan, err := fmt.Scan(&name1, &age1)
	if err != nil {
		return
	} // 等待用户在命令行输入一个值
	fmt.Println("scan", scan)
	fmt.Println("姓名：", name1)
	fmt.Println("年龄：", age1)
	fmt.Println("end")

	var a, b int
	fmt.Println("请按指定格式输入：")
	fmt.Scanf("%d+%d", &a, &b)
	fmt.Println(a + b)
}
```

![image-20230821163940081](README.assets/image-20230821163940081.png)

### 11.流程控制语句

```go
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
```

### 12.循环语句

```go
package main

import "fmt"

func main() {
	/*
	   for 表达式 {
	      //循环语句
	   }
	*/
	var count = 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	for count > 0 {
		fmt.Println(count)
		count--
	}

	// 三要素for循环
	/*
	   var count = 0  //初始语句
	   for count < 10 { // 条件表达式
	       fmt.Println(count)
	       count++ //步进语句
	   }
	*/

	/*
	   for num := 0; num < 10; num++ {
	     // 循环语句
	     fmt.Println(num)
	   }
	*/

	// 基于循环  1+2+3+4+...+100 ?

	var total = 0
	for num := 1; num <= 100; num++ {
		total += num
	}
	fmt.Println("total", total)
}
```

### 13.循环语句和分支语句之间的嵌套

```go
package main

import "fmt"

func main() {

	// 循环语句中嵌套分支语句
	for num := 1; num <= 100; num++ {
		if num%2 == 0 {
			fmt.Println(num)
		}
	}

	// 分支语句中嵌套循环语句
	var count int
	fmt.Println("请输入一个正整数：")
	fmt.Scan(&count)

	if count > 100 {
		// 从小到大打印 1 - num
		for i := 1; i < count; i++ {
			fmt.Println(i)
		}
	} else {
		// 从大到小打印num - 1
		for i := count; i > 0; i-- {
			fmt.Println(i)
		}
	}
}
```

### 14.for循环作用域

```go
package main

import "fmt"

func foo() {
	var y = 100
	fmt.Println(y)
}

func main() {

	var x = 10
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		fmt.Println(x) //内部可以拿到外部的 x
	}
	//fmt.Println(i)   //拿不到 i
}
```

### 15.退出循环

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {

		if i == 3 {
			continue //退出当次循环
		}

		if i == 6 {
			break //退出循环
		}

		fmt.Println(i)
	}
}
```

## 三.重要数据类型

### 1.指针类型

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// &变量 ： 获取变量地址
	var x = 10
	fmt.Printf("赋值之前x的对应地址%p\n", &x)
	x = 100
	fmt.Printf("赋值之后x的对应地址%p\n", &x)

	var y string
	y = "hello"
	fmt.Println("y", y)

	var p *int // p 是一个整型指针类型
	p = &x     //  var p = &x
	fmt.Println("p", p)

	// *指针变量：通过取值操作 获取地址存的值
	fmt.Println("p这个地址里的值:", *p, ",类型：", reflect.TypeOf(*p))
	*p = 1000
	fmt.Println(x)

	/*
		var a = 1
		var b = a
		b = 100 // 改了b 不会影响 a
		fmt.Println(a, b)
	*/

	var a = 1
	var b = &a
	var c *int
	c = b

	*b = 100 //  a  *b  *c 值相同
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(*c)

	// 案例1
	var a1 = 100
	var b1 = &a1    // a1的地址
	var c1 = &b1    // b1的地址
	**c1 = 200      // *c1 拿到 b1的地址里存的值 即 a1 的地址 **c1拿到 a1地址里的值 然后取修改为 200
	fmt.Println(a1) // 200

	// 案例2
	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)  //2
	fmt.Println(*p2) //2
}
```

### 2.new函数

```go
package main

import "fmt"

func main() {
    // 基本数据类型 （整型，浮点型，字符串，布尔类型）范围属于值类型
    // 值类型特点，声明未赋值之前存在一个默认值（zero value）
    var x int
    fmt.Println(x)

    var name string
    fmt.Println(name)

    // 指针类型属于引用类型，包括切片，map，channel都是引用类型
    // 引用类型当声明未赋值之前是没有开辟空间的，即没有默认值
    var p *int
    p = new(int) // 用new 函数开辟一个空间

    *p = 10
    fmt.Println(p)
}
```

![image-20230822151317910](README.assets/image-20230822151317910.png)

### 3.数组

```go
package main

import "fmt"

func main() {
	// 1. 先声明再赋值
	// 数组的声明
	// 数组必须限制长度
	var arr [3]int
	fmt.Println(arr)

	// 索引赋值
	arr[0] = 10
	fmt.Println(arr)

	// 2.数组的声明并赋值
	var names = [3]string{"Nash", "Bob", "Mike"}
	fmt.Println(names) // [Nash Bob Mike]
	var ages = [3]int{22, 23, 24}
	fmt.Println(ages)

	// 3.省略长度赋值  使用 ... 代替
	var fruits = [...]string{"apple", "peach", "pear"}
	fmt.Println(fruits)
}
```

<img src="README.assets/image-20230822153811509.png" alt="image-20230822153811509" style="zoom:150%;" />

### 4.切片

切片是一个动态数组，因为数组的长度是固定的，所以操作起来很不方便。因此在开发中数组并不常用，切片类型才是大量使用的

