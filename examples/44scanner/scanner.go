package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

// 演示行过滤器
// 在bash命令行中 grep和sed是常见的行过滤器
func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
