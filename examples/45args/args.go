package main

import (
	"fmt"
	"os"
)

// 命令行参数演示示例代码

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

}
