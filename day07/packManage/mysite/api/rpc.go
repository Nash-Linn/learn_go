package api

import (
	"fmt"
	"mysite/db"
)

func RpcApi() {
	db.HandleRedis()
	fmt.Println("操作rpcApi接口")
}
