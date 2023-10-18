package service

import (
	"customerSys/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CustomerService struct {
	Customers  []model.Customer
	CustomerID int
}

func (cs *CustomerService) AddCustomer() {
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

	cs.CustomerID++
	var newCustomer = model.Customer{
		Cid:    cs.CustomerID,
		Name:   name,
		Gender: gender,
		Age:    age,
		Email:  email,
	}
	cs.Customers = append(cs.Customers, newCustomer)
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------添加客户成功------------------
			`)
}

func (cs *CustomerService) ListCustomer() {
	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表开始------------------
			`)

	for _, customer := range cs.Customers {
		fmt.Printf("客户编号：%-6d,姓名：%-8s,性别：%-6s,年龄：%-8d,邮箱：%-8s\n",
			customer.Cid, customer.Name, customer.Gender, customer.Age, customer.Email)
	}

	fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表结束------------------
			`)
}

func (cs *CustomerService) UpdateCustomer() {
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
		for i, customer := range cs.Customers {
			if customer.Cid == cid {
				updateIndex = i
				break
			}
		}
		if updateIndex == -1 {
			fmt.Println("您输入的客户编号不存在")
			continue
		}
		updateCustomer := &cs.Customers[updateIndex]

		var name string
		fmt.Printf("请输入客户姓名(%s)：", updateCustomer.Name)
		fmt.Scanln(&name)

		var gender string
		fmt.Printf("请输入客户性别(%s)：", updateCustomer.Gender)
		fmt.Scanln(&gender)

		var age int
		fmt.Printf("请输入客户年龄(%s)：", updateCustomer.Age)
		fmt.Scanln(&age)

		var email string
		fmt.Printf("请输入客户邮箱(%s)：", updateCustomer.Email)
		fmt.Scanln(&email)

		if name != "" {
			updateCustomer.Name = name
		}
		if gender != "" {
			updateCustomer.Gender = gender
		}
		if age != 0 {
			updateCustomer.Age = age
		}
		if email != "" {
			updateCustomer.Email = email
		}

		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
				------------------客户修改成功------------------
				`)
		break
	}
}

func (cs *CustomerService) DeleteCustomer() {

	for true {
		//引导用户输入一个客户编号
		var cid int
		fmt.Print("请输入删除的客户编号：")
		fmt.Scan(&cid)
		// 判断客户是否存在
		var deleteIndex int = -1
		for i, customer := range cs.Customers {
			if customer.Cid == cid {
				deleteIndex = i
				break
			}
		}
		if deleteIndex == -1 {
			fmt.Println("您输入的客户编号不存在")
			continue
		}
		cs.Customers = append(cs.Customers[:deleteIndex], cs.Customers[deleteIndex+1:]...)
		fmt.Print("删除成功")
		break
	}
}

func (cs *CustomerService) SaveCustomer() {
	customersJsonBytes, _ := json.Marshal(cs.Customers)

	err := ioutil.WriteFile("db/customer.json", customersJsonBytes, 0666)

	if err != nil {
		fmt.Println("err：", err)
	}

	fmt.Println("保存成功!")
}
