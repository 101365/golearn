package main

import "fmt"

// 本示例演示闭包

/**
	关于闭包的解释
	闭包只是形式和表现上像函数，但实际上不是函数(有点绕)。区别在哪里呢？
	函数是可执行的代码，这些代码在函数被定义后就确定了，不会在执行时发送变化，所以可以认为一个函数只有一个实例。
	闭包则不同，它在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例
	go语言中，函数是一阶值，即函数可以作为另一个函数的返回值或者参数，一个函数对象和一个数字是类似的。
	go语言中，函数可以嵌套定义。
 */

func seed() func() int  {
	seed := 0
	return func() int {
		seed += 1
		return seed
	}
}

func main() {

	next := seed()
	// 每次调用结果都不一样，这是因为与该函数绑定的引用环境(即本例中的seed数值)发生了变化
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	// 重新初始化
	newNext := seed()
	fmt.Println(newNext())

}
