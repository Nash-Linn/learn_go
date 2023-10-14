package main

import "fmt"

func main() {
	var customer1 = map[string]string{"name": "Ken", "age": "22"}
	var customer2 = map[string]string{"name": "Nash", "age": "23"}
	var customer3 = map[string]string{"name": "Mary", "age": "24"}

	var customers = []map[string]string{customer1, customer2, customer3}

	// 添加客户
	var newCustomer = map[string]string{"name": "Jack", "age": "25"}
	customers = append(customers, newCustomer)

	fmt.Println(customers)
	// 查看
	for _, customer := range customers {
		fmt.Printf("客户名字：%-8s,年龄：%-8s\n", customer["name"], customer["age"])
	}

	//修改 按索引修改
	customers[1]["name"] = "Nash2"
	fmt.Println(customers)

	//删除 按索引删除
	customers = append(customers[:1], customers[2:]...)
	fmt.Println(customers)

}
