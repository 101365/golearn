package main

import "fmt"

// 本示例重点不是代码，而是理解指针的行为

// 传值传递的是副本
func changeVal(v int) {
	v = 0
}

// 指针作为参数
func changePtr(ptr *int)  {
	*ptr = 0;
}

func main() {

	i := 100
	fmt.Println("original: ", i) // 100
	changeVal(i)
	fmt.Println("after change value: ", i) // 100

	b := 100
	fmt.Println("original b: ", b) // 100
	changePtr(&b)
	fmt.Println("after change pointer: ", b) // 0

}


