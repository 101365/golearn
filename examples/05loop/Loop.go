package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 2; j <= 8; j++ {
		fmt.Println(j)
	}

	// 演示跳出循环体
	for {
		fmt.Println("loop ... ", i)
		i++
		if i > 50 {
			break
		}
	}

	// 演示goto关键字
	var k int
	for {
		fmt.Println("goto ... ", k)
		k++
		if k == 30 {
			goto LABEL1
		}
	}

	LABEL1:
		fmt.Println("goto label1 ...")

}
