package main

import "fmt"

// 演示结构体的定义、新建、访问

type user struct {
	age int
	name string
}

func main() {

	// 根据字段定义顺序依次赋值
	u := user{18, "larry"}
	fmt.Println(u)
	// 由于是在同一个包内 可以访问首字母小写的字段
	fmt.Println(u.name)

	// 指定字段创建
	u2 := user{name:"larry"}
	fmt.Println(u2)
	fmt.Println(u2.name)

	up := &u
	fmt.Println(up.name)

	// 通过指针修改属性
	up.name = "lisi"
	fmt.Println(u.name)

}
