package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("------------------go开始运行了-------------")
	args := os.Args[1:] // 获取除了程序名称外的参数
	fmt.Println("命令行参数:")
	for _, arg := range args {
		fmt.Println(arg)
	}
	fmt.Println("------------------go结束运行了-------------")
}
