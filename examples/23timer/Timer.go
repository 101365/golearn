package main

import (
	"time"
	"fmt"
)

// 演示定时器
// 定时器的作用是在未来的某个时间节点执行一次
func main() {

	t1 := time.NewTimer(time.Second * 2)
	fmt.Println(<-t1.C)

	fmt.Println("Timer 1 expired")

	t2 := time.NewTimer(time.Second)

	go func() {
		<- t2.C
		fmt.Println("Timer 2 expired")
	}()

	// 和time.sleep相比，定时器可以取消
	stop := t2.Stop()
	if stop {
		fmt.Println("Timer 2 stoped")
	}

}
