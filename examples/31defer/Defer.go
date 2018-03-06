package main

import (
	"fmt"
	"os"
)

// 演示defer的使用
// 类比Java，相当于finally的作用，用来确保一个函数调用在程序执行结束前执行
// 在Java中有个经典的问题，即finally配合return时，返回值是多少的问题

func main() {

	f := createFile("/tmp/demo.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
