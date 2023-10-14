package main

import (
	"encoding/json"
	"fmt"
)

type Addr struct {
	Province string
	City     string
}

type Stu struct {
	Name string `json:"name"` //结构体的标签
	Age  int    `json:"-"`    //表示不参与序列化
	Addr Addr

	//gender  string //首字母小写的字段不会被序列化
}

func main() {
	var nashMap = map[string]interface{}{
		"name": "Nash",
		"age":  18,
		"addr": map[string]string{
			"province": "浙江",
			"city":     "杭州",
		},
	}

	var nashStruct = Stu{Name: "Nash", Age: 18, Addr: Addr{Province: "浙江", City: "杭州"}}

	jsonNashMap, _ := json.Marshal(nashMap)
	fmt.Println(string(jsonNashMap))

	jsonNashStruct, _ := json.Marshal(nashStruct)
	fmt.Println(string(jsonNashStruct))

	// 反序列化
	var nashMapUnJson map[string]interface{}
	json.Unmarshal(jsonNashMap, &nashMapUnJson)
	fmt.Println("nashMapUnJson", nashMapUnJson)

	var nashStructUnJson Stu
	json.Unmarshal(jsonNashStruct, &nashStructUnJson)
	fmt.Println("nashStructUnJson", nashStructUnJson)

	fmt.Println("--------------------------------------")

	var s1 = Stu{Name: "S1", Age: 18, Addr: Addr{Province: "浙江", City: "杭州"}}
	var s2 = Stu{Name: "S2", Age: 18, Addr: Addr{Province: "浙江", City: "杭州"}}
	var s3 = Stu{Name: "S3", Age: 18, Addr: Addr{Province: "浙江", City: "杭州"}}

	var data = []Stu{s1, s2, s3}

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))

	var dataUnJson []Stu
	json.Unmarshal(jsonData, &dataUnJson)
	fmt.Println(dataUnJson)
}
