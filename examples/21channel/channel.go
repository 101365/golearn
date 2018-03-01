package main

import (
	"fmt"
	"time"
)

// 演示通道

func main() {

	// channel <- 语法 发送一个新的值到通道中
	// <-channel 语法从通道中接收一个值

	list := make(chan int)

	go func() {
		list <- 12
	}()

	var value int
	value = <-list
	fmt.Println(value)

	// 缓冲通道：可指定最多缓冲多少个值
	infos := make(chan int, 2)
	infos <- 73
	infos <- 56
	//infos <- 89

	fmt.Println(<-infos)
	fmt.Println(<-infos)

	// 通道作为协程的通讯信号
	flag := make(chan bool, 1)
	worker(flag)

	<-flag

	// 演示通道方向
	showDirection()

}

// 演示通道方向
func showDirection() {
	// 一般在定义函数参数时使用
}

// 工作协程
func worker(flag chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	flag <- true
}
