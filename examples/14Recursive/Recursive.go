package main

import "fmt"

// 演示递归 和其他语言一致 这里演示阶乘的实现

func factorial(v int) int  {
	if v == 1 {
		return 1
	}
	return v * factorial(v - 1)
}

func main() {

	total := factorial(5)
	fmt.Println(total)

}
