package main

import (
	"io/ioutil"
	"fmt"
	"io"
	"bufio"
	"os"
)

// 演示文件的读取 写入 读文件 写文件


func main() {

	// 文件读取示例
	readFileDemo()

	// 文件写入示例
	writeFileDemo()

}

func writeFileDemo() {

	d1 := []byte("hello\ngo\n")
	// 第三个参数是权限
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// 创建文件
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}

// 演示文件读取
func readFileDemo() {

	// 读取文件
	dat, err := ioutil.ReadFile("/tmp/dat")
	// 检查错误
	check(err)
	fmt.Print(string(dat))

	// 打开文件
	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close()

	// 创建一个切片
	b1 := make([]byte, 5)
	// 读取到切片
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 从已知位置开始读取，6表示offset，0代表从文件开始计算offset
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	//
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// 最少读取多少个字节 当不足时会报错
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	// 少于指定数量的字节时会报错
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))



}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
