package main

/*
 在Go语言中，结构体承担着面向对象语言中类的作用。Go语言中，结构体本身仅用来定义属性。还可以通过接收器函数来定义方法，
 使用内嵌结构体来定义继承，这样使用结构体相关操作Go语言就可以实现OOP面向对象编程
*/

// 类型包含属性和方法
// 对象是类型实例化的结果

// 声明结构体  关键字  type ... struct
type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

type Person struct {
	name   string
	gender string
	age    int8
	addr   string
}

func main() {

}
