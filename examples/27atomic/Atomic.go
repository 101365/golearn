package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 演示原子计数器
// 并发模型，大致有2种，一种是CSP(即通过通道+协程的方式)，一种是内存共享。
// go语言中这两种都可以支持。
// 本示例中，先演示一个不安全的示例，然后演示一个通过原子计数修正为安全的。
//
func main() {

	// 演示不安全的计数器
	unsafeDemo()

	// 演示原子计数器
	safeDemo()

}

func safeDemo() {
	num := 500
	var ops uint64 = 0

	var group sync.WaitGroup
	group.Add(num)

	// 启动500个协程，来完成递增操作
	for i:=0; i<num; i++ {
		go func() {
			defer group.Done()
			atomic.AddUint64(&ops, 1)
		}()
	}

	// 等待所有的协程都执行完毕
	group.Wait()

	// 结果是500
	fmt.Println("atomic counter => ", ops)
}

// 该示例不保证百分之百复现错误，多运行几次就会出现ops最终不是500的情况
func unsafeDemo() {
	num := 500
	var ops uint64 = 0

	var group sync.WaitGroup
	group.Add(num)

	// 启动500个协程，来完成递增操作
	for i:=0; i<num; i++ {
		go func() {
			defer group.Done()
			ops++
		}()
	}

	// 等待所有的协程都执行完毕
	group.Wait()

	// 上面执行了500次递增操作 值为500才是正确的  这里因为有竞态条件 会出现错误结果
	fmt.Println("unsafe counter => ", ops)

}
