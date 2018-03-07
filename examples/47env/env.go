package main

import (
	"fmt"
	"strings"
	"os"
)

// 环境变量示例代码  类似于Java的System.property
func main() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// 使用 os.Environ 来列出所有环境变量键值队。
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

}
