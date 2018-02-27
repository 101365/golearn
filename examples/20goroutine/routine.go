package main

import (
	"fmt"
)

// 演示协程示例
/*
	协程是轻量级的线程
 */

func master(s string)  {
	for i:=0; i<3; i++ {
		fmt.Println(s, ":", i)
	}
}

func main() {

	master("main")

	go master("routine")

	go func(msg string) {
		fmt.Println(msg)
	}("haha")

	// 需要将主协程block住 这里需要按下任意键才会停止
	var k int
	fmt.Scanln(&k)
	fmt.Println("done")

}
