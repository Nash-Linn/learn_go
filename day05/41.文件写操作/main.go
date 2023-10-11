package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/*
		OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
		如果操作成功，返回的文件对象可用于I/O；如果出错，错误底层类型是*PathError。
	*/

	file, err := os.OpenFile("满江红", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	fmt.Println("file:", file)

	// (1) 按字符串或者字节写操作
	file.WriteString("怒发冲冠，凭栏处、潇潇雨歇。抬望眼、仰天长啸，壮怀激烈。\n")

	file.Write([]byte("三十功名尘与土，八千里路云和月。\n"))

	// (2) 缓存写
	// 创建一个写入器（缓存区）
	writer := bufio.NewWriter(file)
	// 将数据写入缓存
	writer.WriteString("莫等闲，白了少年头，空悲切。\n")
	// 将缓存中的内容写入文件
	writer.Flush()

	// （3）写整个文件
	str := "靖康耻，犹未雪；臣子恨，何时灭？\n"
	ioutil.WriteFile("满江红1", []byte(str), 0666)
}
