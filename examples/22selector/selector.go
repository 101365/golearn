package main

import (
	"time"
	"fmt"
)

// 通道选择器

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	for i := 0; i<2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(i, "get", msg1)
		case msg2 := <-c2:
			fmt.Println(i, "get", msg2)
		}
	}



}
