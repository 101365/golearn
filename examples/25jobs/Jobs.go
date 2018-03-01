package main

import (
	"fmt"
	"time"
)

/**
	模拟一个工作池。
	业务逻辑：
	初始化若干个任务，有多个worker领取任务并执行，最后收集任务结果。
 */

func worker(id int, jobs <-chan int, results chan<- int)  {

	for job := range jobs {
		fmt.Println(id, " start job ", job)
		time.Sleep(time.Second * 1)
		results <- job * 2
	}

}

func main() {

	length := 10

	jobs := make(chan int, length)
	results := make(chan int, length)

	for i:=0; i<3; i++ {
		go worker(i + 1, jobs, results)
	}

	for i:=0; i<length; i++ {
		jobs <- i
	}

	// 关闭
	close(jobs)
	fmt.Println("jobs closed")

	// 打印结果集 这里注意 因为没有合适的地方关闭results通道，为避免deadlock只能通过设定长度检测
	for i:=0; i<length; i++ {
		fmt.Println(<-results)
	}
}
