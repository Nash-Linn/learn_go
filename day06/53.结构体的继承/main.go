package main

import "fmt"

// Animal 类型
type Animal struct {
	Name string
}

func (a Animal) eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

// Dog 类型
type Dog struct {
	Kind string
	Animal
}

func (d Dog) bark() {
	fmt.Printf("%s is barking\n", d.Name)
}

// Cat 类型
type Cat struct {
	Kind string
	Animal
}

func main() {
	d1 := Dog{
		Kind: "哈士奇",
		Animal: Animal{
			Name: "旺财",
		},
	}
	d1.eat()
	d1.bark()

	c1 := Cat{
		Kind: "波斯猫",
		Animal: Animal{
			Name: "小花",
		},
	}
	c1.eat()
}
