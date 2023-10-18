package api

import (
	"fmt"
	"learn/day07/packManage/mysite/db"
)

func RestfulApi() {
	db.HandleMySql()

	fmt.Println("操作restfulApi接口")
}
