package main

import "os"

// 演示panic

// panic 意味着有些出乎意料的错误发生。
// 通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误。
// panic 的一个基本用法就是在一个函数返回了错误值但是我们并不知道如何（或者不想）处理时终止运行。
// 这里是一个在打开一个新文件时返回异常错误时的panic 用法。
func main() {

	//panic("a problem")

	_, err := os.Open("/tmp/somefile")
	if err != nil {
		panic(err)
	}

}
