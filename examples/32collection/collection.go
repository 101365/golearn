package main

import (
	"fmt"
	"strings"
)

// 演示集合函数
// 在日常开发过程中，我们经常会涉及到对集合中数据的操作，比如：
// A.选择所有的满足某个条件的元素
// B.将所有元素映射到一个新的集合中

// 在其他语言中，惯用的方式是使用泛型。在go语言中，没有泛型机制，对应的方案是在需要的时候提供合适的集合函数来满足要求。

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}

// 类似于Java的string.indexOf(str)
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}


// 是否包含某个字符串
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 只要有一个满足断言函数就返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 全部字符串都满足断言函数时才返回true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 满足条件的过滤到新切片
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 映射到新切片
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
