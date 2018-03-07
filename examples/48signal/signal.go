package main

import (
	"os/signal"
	"syscall"
	"fmt"
	"os"
)

// 信号

func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 相当于订阅指定的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 接收到对应的信号后 程序结束
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
