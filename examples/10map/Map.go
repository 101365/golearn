package main

import (
	"fmt"
)

// map的使用
func main() {

	// 示例一
	m := make(map[string]int)
	m["a"] = 12
	m["b"] = 6
	fmt.Println(m)

	// 删除某个键
	delete(m, "b")
	fmt.Println(m)

	// 直接初始化map
	n := map[string]int{"foo":1, "hi":3}
	fmt.Println(n)

	// 遍历map
	for k, v := range n {
		fmt.Println(k, " = ", v)
	}

	// 判断某个键是否存在
	value, exists := n["kkk"]
	fmt.Println(value)
	fmt.Println(exists)
	if exists {
		fmt.Println(value)
	}

}
