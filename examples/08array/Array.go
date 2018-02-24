package main

import "fmt"

// 本示例演示数组的常见定义
func main() {

	/**
		数组是固定长度、存放固定类型数据的结构

	 */

	 // 示例1：先声明，后设置数值
	var x [3]int
	fmt.Println("empty array => ", x)
	x[1] = 100
	fmt.Println("one element array => ", x)
	fmt.Println("get x[1] => ", x[1])

	// 示例二: 指定长度直接初始化数组
	b := [3]int{1, 2, 3}
	fmt.Println(b)
	fmt.Println("b length => ", len(b))

	// 示例三：不指定长度直接初始化，自动计算长度(注意使用...)
	c := [...]int{1, 2}
	fmt.Println(c)
	fmt.Println("c length => ", len(c))

	// 示例四：多维数组
	var ka [2][3]int
	for i:=0; i<2; i++ {
		for j:=0; j<3; j++ {
			ka[i][j] = i + j
		}
	}

	fmt.Println("2d array => ", ka)

}
