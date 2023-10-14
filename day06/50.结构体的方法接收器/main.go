package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func NewStudent(sid int, name string, age int, course []string) Student {
	return Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}
}

// Student类型的方法接收器
// 在方法名前加括号，括号内写结构体的名称，表示这个方法和这个结构体绑定
func (s Student) read(bookName string) {
	fmt.Printf("学生正在读%s书\n", bookName)
}

func (s Student) learn() {
	fmt.Printf("学生%s正在学习\n", s.name)
}

func main() {
	s1 := NewStudent(1001, "Bob", 11, []string{"GO"})
	s1.read("数学")
	s1.learn()
}
