package main

import (
	"fmt"
)

// 切片示例代码
func main() {

	// 简单起见 本例中所有的切片都是int类型

	// 示例一：使用make函数生成
	s := make([]int, 3)
	fmt.Println(s)
	fmt.Println(len(s), " = ", cap(s))

	// 示例二：指定len和cap容量生成
	ss := make([]int, 2, 4)
	fmt.Println(ss)
	fmt.Println(len(ss), " = ", cap(ss))

	// 示例三：通过数组生成
	arr := [6]int{1, 9, 34, 23, 16}
	fmt.Println(arr)
	s3 := arr[3:5]
	fmt.Println(s3)

	// 示例四：提前定义
	var s4 []int
	fmt.Println(len(s4)) // 此时长度为0
	s4 = append(s4, 12)
	fmt.Println(len(s4))

	// 示例五：直接定义初始化
	s5 := []int{1,2,3}
	fmt.Println(len(s5))

	// 切片相关操作

	// 拷贝
	src := []int{1,2,3}
	dest := make([]int, 6)
	copy(dest, src)
	fmt.Println(dest)

	// 截取
	d1 := []int{2, 5, 78, 66, 123, 45}
	d2 := d1[:3]
	fmt.Println(d2) // [2, 5, 78]
	d3 := d1[3:]
	fmt.Println(d3) // [66 123 45]
	d4 := d1[2:4]
	fmt.Println(d4) // [28, 66]


}
