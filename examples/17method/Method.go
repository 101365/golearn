package main

import (
	"fmt"
)

// 演示方法

// 方法和函数很像，可以简单理解：方法是指定了接收器的函数。
// 方法的接收器只能是当前文件中定义的类型

type User struct {
	name string
}

// 值类型接收器
func (a User) add(x, y int) int  {
	return x + y
}

// 指针类型接收器
func (a *User) concat(x, y string) string  {
	return x + " " + y
}

func main() {

	// 至于值类型和指针类型接收器的区别 本示例不做深入说明

	user := User{"larry"}

	sum := user.add(1, 2)
	fmt.Println(sum)

	result := user.concat("hello", "world")
	fmt.Println(result)

}