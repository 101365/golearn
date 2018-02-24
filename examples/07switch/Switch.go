package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println(i)

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("今天是星期六")
	default:
		fmt.Println("今天不是星期六")
	}

	var now = time.Now().Weekday()
	fmt.Println(now)

	t := time.Now()
	fmt.Println(t.Minute())

	// 注意和上面示例的区别 这里的条件可以使用任意的变量
	switch {
	case t.Minute() < 30 :
		fmt.Println("上半时")
	case i == 2:
		fmt.Println("i == 2")
	default:
		fmt.Println("下半时")
	}

}
