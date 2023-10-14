package main

import (
	"fmt"
	"os"
)

/*

 面向过程 面向对象 对比

 例子
 面向过程：
 1. 将衣服放进洗衣机
 2. 倒入洗衣液
 3. 洗涤
 4. 脱水
 5. 甩干
 6. 晾晒

 面向对象：
 人对象-人类型
 洗衣机-洗衣机类型

 1. 人.将衣服放进洗衣机()
 2. 人.倒入洗衣液()
 3. 洗衣机.洗涤()
 4. 洗衣机.脱水()
 5. 洗衣机.甩干()
 6. 人.晾晒()

*/

type Customer struct {
	cid    int
	name   string
	gender string
	age    int
	email  string
}

type CustomerService struct {
	customers  []Customer
	customerID int
}

//var customers []Customer
//var customerID int

func (cs *CustomerService) addCustomer() {
	// 引导用户输入

	//引导用户输入学号和姓名
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------添加客户开始------------------
			`)
	var name string
	fmt.Print("请输入客户姓名：")
	fmt.Scan(&name)

	var gender string
	fmt.Print("请输入客户性别：")
	fmt.Scan(&gender)

	var age int
	fmt.Print("请输入客户年龄：")
	fmt.Scan(&age)

	var email string
	fmt.Print("请输入客户邮箱：")
	fmt.Scan(&email)

	cs.customerID++
	var newCustomer = Customer{
		cid:    cs.customerID,
		name:   name,
		gender: gender,
		age:    age,
		email:  email,
	}
	cs.customers = append(cs.customers, newCustomer)
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------添加客户成功------------------
			`)
}

func (cs *CustomerService) listCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表开始------------------
			`)

	for _, customer := range cs.customers {
		fmt.Printf("客户编号：%-6d,姓名：%-8s,性别：%-6s,年龄：%-8d,邮箱：%-8s\n",
			customer.cid, customer.name, customer.gender, customer.age, customer.email)
	}

	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表结束------------------
			`)
}

func (cs *CustomerService) updateCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
				------------------客户修改开始------------------
				`)
	for true {
		//引导用户输入一个客户编号
		var cid int
		fmt.Print("请输入修改的客户编号：")
		fmt.Scan(&cid)

		// 判断客户是否存在
		var updateIndex int = -1
		for i, customer := range cs.customers {
			if customer.cid == cid {
				updateIndex = i
				break
			}
		}
		if updateIndex == -1 {
			fmt.Println("您输入的客户编号不存在")
			continue
		}
		updateCustomer := &cs.customers[updateIndex]

		var name string
		fmt.Printf("请输入客户姓名(%s)：", updateCustomer.name)
		fmt.Scanln(&name)

		var gender string
		fmt.Printf("请输入客户性别(%s)：", updateCustomer.gender)
		fmt.Scanln(&gender)

		var age int
		fmt.Printf("请输入客户年龄(%s)：", updateCustomer.age)
		fmt.Scanln(&age)

		var email string
		fmt.Printf("请输入客户邮箱(%s)：", updateCustomer.email)
		fmt.Scanln(&email)

		if name != "" {
			updateCustomer.name = name
		}
		if gender != "" {
			updateCustomer.gender = gender
		}
		if age != 0 {
			updateCustomer.age = age
		}
		if email != "" {
			updateCustomer.email = email
		}

		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
				------------------客户修改成功------------------
				`)
		break
	}
}

func (cs *CustomerService) deleteCustomer() {

	for true {
		//引导用户输入一个客户编号
		var cid int
		fmt.Print("请输入删除的客户编号：")
		fmt.Scan(&cid)
		// 判断客户是否存在
		var deleteIndex int = -1
		for i, customer := range cs.customers {
			if customer.cid == cid {
				deleteIndex = i
				break
			}
		}
		if deleteIndex == -1 {
			fmt.Println("您输入的客户编号不存在")
			continue
		}
		cs.customers = append(cs.customers[:deleteIndex], cs.customers[deleteIndex+1:]...)
		fmt.Print("删除成功")
		break
	}
}

func main() {
	cs := CustomerService{
		customers:  []Customer{},
		customerID: 0,
	}

	for true {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
	------------------客户信息管理系统------------------
	1.添加客户                                         
	2.查看客户                                         
	3.更新客户
	4.删除客户
	5.保存
	6.退出
	--------------------------------------------------
`)
		var choice int8

		fmt.Print("请输入您的选择：")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// 添加客户
			cs.addCustomer()
		case 2:
			// 查看
			cs.listCustomer()
		case 3:
			//修改
			cs.updateCustomer()

		case 4:
			//删除
			cs.deleteCustomer()

		case 5:
			//保存
			fmt.Println("保存成功")
		case 6:
			//退出
			fmt.Println("退出成功")
			os.Exit(0)
		}
	}
}
