package main

import (
	"time"
	"fmt"
)

// 演示速率限制示例

func main() {

	fmt.Println("start common limit ...")
	// 全局限速示例
	limitExample()

	fmt.Println("=====================")

	fmt.Println("start burst limit ...")
	// 脉冲限速示例
	burstyLimiterExample()

}

// 脉冲限速示例
func burstyLimiterExample() {

	/*
		特定场景，允许段脉冲串请求，同时保持总体速率限制。
		表现为前三个立刻提供服务，后面的限速200毫秒。
	 */

	burstyLimiter := make(chan time.Time, 5)
	for i:=0; i<3; i++ {
		burstyLimiter <- time.Now()
	}

	// 启动一个协程
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 模拟五个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}

// 全局限速示例
func limitExample() {

	/*
		该示例演示200毫秒处理一个请求
	 */
	 requests := make(chan int, 5)

	 // 初始化5个请求
	 for i:=1; i<=5; i++ {
	 	requests <- i
	 }
	 close(requests)

	 // 速率限制在200毫秒
	 limiter := time.Tick(time.Millisecond * 200)
	 for req := range requests {
	 	<- limiter
	 	fmt.Println("request", req, time.Now())
	 }

}
