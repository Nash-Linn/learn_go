package api

import (
	"fmt"
	"mysite/db"
)

func RestfulApi() {
	db.HandleMySql()

	fmt.Println("操作restfulApi接口")
}
