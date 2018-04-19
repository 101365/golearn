package main

import "testing"

// 验证并发场景下，map赋值是否为原子操作

var m = make(map[string]string)

func TestMultiReadAndWrite(t *testing.T) {

	for i:=0; i<100000; i++ {
		// 模拟写
		go func() {

		}()

		// 模拟读
		go func() {

		}()
	}

}

func write()  {
	tmp := make(map[string]string)
	tmp["s"] = "s"

	// 赋值
	m = tmp
}

func read()  {

}




