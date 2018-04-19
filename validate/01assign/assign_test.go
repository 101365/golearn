package main

import (
	"testing"
	"fmt"
	"sync"
	"runtime"
	"encoding/json"
)


func TestMultiReadAndWrite(t *testing.T) {

	runtime.GOMAXPROCS(runtime.NumCPU())

	m := write(7)

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	loop := 100
	var wg sync.WaitGroup
	wg.Add(loop)

	// 启动若干个进程同时访问
	for i:=0; i<loop; i++ {
		go func(index int) {

			defer func() {
				wg.Done()
			}()

			reqJson, _ := json.Marshal(m)

			fmt.Println(reqJson)

		}(i)
	}

	// 等待执行完成
	wg.Wait()

	fmt.Println("over ...")

}

func write(size int) *map[string]string {
	tmp := make(map[string]string)
	for i:=0; i<size; i++ {
		k := fmt.Sprintf("%v", i)
		tmp["key_" + k] = "value_" + k
	}
	fmt.Println(tmp)
	return &tmp
}





