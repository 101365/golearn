package main

import (
	"fmt"
	s "strings"
)

// 演示字符串的常用操作
// 这些都是标准库strings里的函数

var p = fmt.Println

func main() {

	// 是否包含
	p("Contains:  ", s.Contains("test", "es")) // Contains:   true
	// 包含的个数
	p("Count:     ", s.Count("test", "t")) // Count:      2
	// 前缀 startWith
	p("HasPrefix: ", s.HasPrefix("test", "te")) // HasPrefix:  true
	// 后缀 endWith
	p("HasSuffix: ", s.HasSuffix("test", "st")) // HasSuffix:  true
	// indexOf
	p("Index:     ", s.Index("test", "e")) // Index:      1
	// concat 字符串连接函数
	p("Join:      ", s.Join([]string{"a", "b"}, "-")) // Join:       a-b
	// 这个有点意思
	p("Repeat:    ", s.Repeat("a", 5)) // aaaaa
	// 替换
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // Replace:    f00
	// 替换
	p("Replace:   ", s.Replace("foo", "o", "0", 1)) // Replace:    f0o
	// 切割
	p("Split:     ", s.Split("a-b-c-d-e", "-")) // Split:      [a b c d e]
	// 大写
	p("ToLower:   ", s.ToLower("TEST")) // ToLower:    test
	// 小写
	p("ToUpper:   ", s.ToUpper("test")) // ToUpper:    TEST
	p()


	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
