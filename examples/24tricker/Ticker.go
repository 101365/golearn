package main

import (
	"time"
	"fmt"
)

// 演示打点器 和定时器的区别在于，它会以固定的时间间隔重复执行
func main() {

	// 每个200毫秒定时执行某个逻辑
	t := time.NewTicker(time.Millisecond * 200)

	go func() {
		for c := range t.C {
			fmt.Println("ticker at ", c)
		}
	}()

	time.Sleep(time.Second * 3)
	t.Stop()
	fmt.Println("ticker stopped")

}
