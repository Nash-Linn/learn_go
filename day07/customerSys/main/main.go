package main

import (
	"customerSys/model"
	"customerSys/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	customersJsonBytes, err := ioutil.ReadFile("db/customer.json")

	var Customers []model.Customer
	var CustomerID int
	if err != nil {
		fmt.Println("err：", err)

		Customers = make([]model.Customer, 0, 10)
		CustomerID = 0
	} else {
		json.Unmarshal(customersJsonBytes, &Customers)
		CustomerID = Customers[len(Customers)-1].Cid
	}

	cs := service.CustomerService{
		Customers:  Customers,
		CustomerID: CustomerID,
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
			cs.AddCustomer()
		case 2:
			// 查看
			cs.ListCustomer()
		case 3:
			//修改
			cs.UpdateCustomer()

		case 4:
			//删除
			cs.DeleteCustomer()

		case 5:
			//保存
			cs.SaveCustomer()
		case 6:
			//退出
			fmt.Println("退出成功")
			os.Exit(0)
		}
	}
}
