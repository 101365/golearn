package main

import "fmt"

// 演示变量的多种定义方式

func main() {

	// var 可以声明一个或多个变量
	/**
		可以自动推断已经初始化的变量类型
		当仅仅声明变量，没有给定初始值时，变量会被初始化为零值。不同类型的零值不同。
	 */

	var a string = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// 自动推断类型
	f := "hello"
	fmt.Println(f)

	var g, h = 1, "hello"
	fmt.Println(g, h)

}
