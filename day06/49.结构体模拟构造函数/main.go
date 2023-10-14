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

func main() {
	s1 := NewStudent(1001, "Bob", 11, []string{"GO"})
	fmt.Println(s1)
}
