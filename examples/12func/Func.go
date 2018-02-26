package main

import "fmt"

// 演示函数的定义和使用

func add(x, y int) int {
	return x + y
}

// 多返回值 常见的是返回结果+错误信息
func demo() (int, int)  {
	return 2, 8
}

func main() {

	sum := add(2, 8)

	fmt.Println("2 + 8 = ", sum)

	// 多返回值
	x, y := demo()
	fmt.Println(x, ", ", y)

	// 忽略部分返回值
	_, k := demo()
	fmt.Println(k)

}
