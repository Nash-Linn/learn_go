package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var s = []int{1, 2, 3, 4}

	var m = map[string]string{
		"name": "Nash",
		"age":  "18",
	}

	// json序列化
	//data, _ := json.Marshal(s)
	data, _ := json.Marshal(m)

	fmt.Println("s", s)

	fmt.Println(string(data))

	fmt.Printf("%#v\n", string(data))

	ioutil.WriteFile("data.json", data, 0666)

	// 反序列化
	readData, _ := ioutil.ReadFile("data.json")
	fmt.Println("readData", readData)
	fmt.Println("readDataStr", string(readData))

	readDataTrans := make(map[string]string)

	json.Unmarshal(readData, &readDataTrans)

	fmt.Println("readDataTrans[name]", readDataTrans["name"])
}
