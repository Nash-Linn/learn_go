package main

import "fmt"

func main() {
	// map 嵌套 slice
	var data = make(map[string][]string)
	data["ZJ"] = []string{"杭州", "温州", "嘉兴"}
	data["BJ"] = []string{"朝阳", "海定", "昌平"}
	fmt.Println(data)

	// 查询 浙江的第二个城市
	fmt.Println(data["ZJ"][1])

	fmt.Println("---------------------------------------")

	// 遍历每一个省份以及对应的城市名
	for proStr, citys := range data {
		fmt.Println(proStr)
		for i, v := range citys {
			fmt.Printf("%d.%s", i, v)
		}
		fmt.Println("\n")
	}
}
