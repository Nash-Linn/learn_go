package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// （1）按字节读取
func readByBytes(file *os.File) {
	data := make([]byte, 9)
	file.Read(data)
	fmt.Println(data)
}

// （2）按行读数据
func readByLine(file *os.File) {
	reader := bufio.NewReader(file)

	for true {
		//lineContent, _, _ := reader.ReadLine()
		//fmt.Println(string(lineContent))

		lineContent, err := reader.ReadString('\n')
		fmt.Println(lineContent)
		if err == io.EOF {
			break
		}
	}
}

// （3）读取整个文件
func readByFile() {
	content, _ := ioutil.ReadFile("./满江红")
	fmt.Println(string(content))
}

func main() {
	//os.Open()函数能够打开一个文件，返回一个 *File 和一个 err

	//打开文件
	file, err := os.Open("./满江红")
	if err != nil {
		fmt.Println("打开文件失败")
	} else {
		fmt.Println(file)
	}

	//需要在终端 使用命令 go run main.go 启动

	// （1）按字节读取
	//readByBytes(file)

	// （2）按行读数据
	//readByLine(file)

	// （3）读取整个文件
	readByFile()
}
