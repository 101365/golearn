package main

import (
	"regexp"
	"fmt"
	"bytes"
)

// go语言中内置了对正则表达式的支持

func main() {

	// 是否匹配
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 编译模式
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 模式是否匹配字符串
	fmt.Println(r.MatchString("peach")) // true

	// 模式中查找 => 返回左侧最匹配的字符串
	fmt.Println(r.FindString("peach punch")) //

	// 查找匹配到的索引位置
	fmt.Println(r.FindStringIndex("peach punch"))

	// 类似于group
	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	// 类似于Compile，但是只有一个返回值
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	// 对匹配到的字符串，使用指定的函数来操作
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
