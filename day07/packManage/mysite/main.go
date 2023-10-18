package main // 声明包名

//import (
//	"fmt"
//	"learn/day07/packManage/mysite/api"
//) // 导入依赖包

import (
	"fmt"
	"github.com/jinzhu/now"
)
import . "mysite/api" //也可使用 . 代表导入包内所有函数

func main() {
	RestfulApi()
	RpcApi()
	fmt.Println(now.BeginningOfDay())
	fmt.Scanln()
}
