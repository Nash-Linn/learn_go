package main

import (
	"fmt"
	"os"
	"strconv"
)

var customers = []map[string]string{}
var customerID int

func main() {
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

			var age string
			fmt.Print("请输入客户年龄：")
			fmt.Scan(&age)

			var email string
			fmt.Print("请输入客户邮箱：")
			fmt.Scan(&email)

			customerID++
			var newCustomer = map[string]string{
				"cid":    strconv.Itoa(customerID),
				"name":   name,
				"gender": gender,
				"age":    age,
				"email":  email,
			}
			customers = append(customers, newCustomer)
			fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------添加客户成功------------------
			`)

		case 2:
			// 查看
			fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表开始------------------
			`)

			for _, customer := range customers {
				fmt.Printf("客户编号：%-6s,姓名：%-8s,性别：%-6s,年龄：%-8s,邮箱：%-8s\n",
					customer["cid"], customer["name"], customer["gender"], customer["age"], customer["email"])
			}

			fmt.Printf("\033[1;30;42m%s\033[0m\n", `
			------------------客户列表结束------------------
			`)

		case 3:
			fmt.Printf("\033[1;30;42m%s\033[0m\n", `
				------------------客户修改开始------------------
				`)
			for true {
				//引导用户输入一个客户编号
				var cid string
				fmt.Print("请输入修改的客户编号：")
				fmt.Scan(&cid)

				// 判断客户是否存在
				var updateIndex int = -1
				for i, customer := range customers {
					if customer["cid"] == cid {
						updateIndex = i
						break
					}
				}
				if updateIndex == -1 {
					fmt.Println("您输入的客户编号不存在")
					continue
				}
				updateCustomer := customers[updateIndex]

				var name string
				fmt.Printf("请输入客户姓名(%s)：", updateCustomer["name"])
				fmt.Scanln(&name)

				var gender string
				fmt.Printf("请输入客户性别(%s)：", updateCustomer["gender"])
				fmt.Scanln(&gender)

				var age string
				fmt.Printf("请输入客户年龄(%s)：", updateCustomer["age"])
				fmt.Scanln(&age)

				var email string
				fmt.Printf("请输入客户邮箱(%s)：", updateCustomer["email"])
				fmt.Scanln(&email)

				if name != "" {
					updateCustomer["name"] = name
				}
				if gender != "" {
					updateCustomer["gender"] = gender
				}
				if age != "" {
					updateCustomer["age"] = age
				}
				if email != "" {
					updateCustomer["email"] = email
				}

				fmt.Printf("\033[1;30;42m%s\033[0m\n", `
				------------------客户修改成功------------------
				`)
				break
			}

		case 4:
			for true {
				//引导用户输入一个客户编号
				var cid string
				fmt.Print("请输入删除的客户编号：")
				fmt.Scan(&cid)
				// 判断客户是否存在
				var deleteIndex int = -1
				for i, customer := range customers {
					if customer["cid"] == cid {
						deleteIndex = i
						break
					}
				}
				if deleteIndex == -1 {
					fmt.Println("您输入的客户编号不存在")
					continue
				}
				customers = append(customers[:deleteIndex], customers[deleteIndex+1:]...)
				fmt.Print("删除成功")
				break
			}

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
