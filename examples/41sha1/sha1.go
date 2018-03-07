package main

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 散列经常用生成二进制文件或者文本块的短标识
// git 版本控制系统大量的使用 SHA1 来标识受版本控制的文件和目录
// Go 在多个 crypto/* 包中实现了一系列散列函数。
// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})
func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	// 入参是额外添加的字符切片，一般不需要
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)

}
