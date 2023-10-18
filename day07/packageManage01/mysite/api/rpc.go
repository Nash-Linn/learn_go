package api

import (
	"fmt"
	"learn/day07/packManage/mysite/db"
)

func RpcApi() {
	db.HandleRedis()
	fmt.Println("操作rpcApi接口")
}
