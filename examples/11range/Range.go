package main

import "fmt"

// 演示range 迭代各种数据结构
//  迭代数据
//  迭代切片
//  迭代map
//  迭代字符串中的Unicode编码

func main() {

	// 数组
	a := [...]int{1, 2, 3}
	for _, num := range a {
		fmt.Println(num)
	}
	fmt.Println("===========")

	// 切片
	b := a[1:2]
	for _, num := range b {
		fmt.Println(num)
	}
	fmt.Println("===========")

	// map
	c := map[string]string{ "a": "hello", "b": "world"}
	for k, v := range c {
		fmt.Println(k, " = ", v)
	}

	fmt.Println("===========")

	// 字符串
	s := "hello world"
	// 第一个返回值是rune的起始字节位置  第二个返回值是rune类型
	for _, ss := range s {
		//
		fmt.Println(string(ss))
	}

}


