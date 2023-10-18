package main // 声明包名

//import (
//	"fmt"
//	"learn/day07/packManage/mysite/api"
//) // 导入依赖包

import (
	"fmt"
)

//import A "learn/day07/packManage/mysite/api"  //可以对包起别名
import . "learn/day07/packManage/mysite/api" //也可使用 . 代表导入包内所有函数

func main() {
	RestfulApi()
	RpcApi()

	fmt.Scanln()
}
