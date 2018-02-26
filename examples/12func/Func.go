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

// 可变参数函数  类似于Java中的可变参数列表
func multi(nums ...int) int  {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("求和：", total)
	return total
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

	// 可变参数列表
	multi(1,2)
	multi(1)
	multi(1, 2, 3, 4, 5)

	// 切片作为参数传入 需要解包
	nums := []int{1,2,3}
	multi(nums...)

}
