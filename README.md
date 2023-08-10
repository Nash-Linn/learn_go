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

