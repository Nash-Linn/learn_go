package main

import (
	"fmt"
)

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func initSid(p *Student) {
	(*p).sid = 100
}

func main() {
	// 实例化方式1
	var Nash Student
	Nash.sid = 1
	Nash.name = "Nash"
	Nash.age = 18

	fmt.Println(Nash.name)

	Nash.course = make([]string, 3)
	Nash.course[0] = "Go"

	fmt.Println(Nash)

	//实例化方式2

	var Ken = Student{
		sid:    2,
		name:   "Ken",
		age:    18,
		course: []string{"Go", "Python"},
	}
	fmt.Println(Ken)

	//实例化方式3
	var Jack = Student{3, "Jack", 18, []string{"Go", "Python"}}
	fmt.Println(Jack)

	//实例化方式4
	var Tom = Ken //值拷贝
	Tom.age = 32
	Tom.name = "Tom"
	fmt.Println(Ken)
	fmt.Println(Tom)

	initSid(&Tom)

	fmt.Println(Tom)

	//实例化方式5
	var p = new(Student)
	p.name = "p"
	fmt.Println(p) //&{0 p 0 []}
}
