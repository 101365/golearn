package main

import "fmt"

// 演示函数的定义和使用

func add(x, y int) int {
	return x + y
}

func main() {

	sum := add(2, 8)

	fmt.Println("2 + 8 = ", sum)

}
