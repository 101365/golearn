package main

import (
	"testing"
	"strings"
	"fmt"
)


func TestCore(t *testing.T) {

	s := "hello world"
	// 创建 Reader
	r := strings.NewReader(s)

	fmt.Println(r) // &amp;{hello world 0 -1}
	fmt.Println(r.Size()) // 11 获取字符串长度
	fmt.Println(r.Len()) // 11 获取未读取长度

	// 读取前6个字符
	for i:=0; i<6; i++ {
		b, err := r.ReadByte() // 读取1 byte
		fmt.Println(string(b), err, r.Len(), r.Size())
	}

	// 读取还未被读取字符串中5字符的数据
	b_s := make([]byte, 5)
	n, err := r.Read(b_s)
	fmt.Println(string(b_s), n ,err) // world 5 &lt;nil&gt;
	fmt.Println(r.Size()) // 11
	fmt.Println(r.Len()) // 0

}