package main

import (
	"testing"
	"strings"
	"fmt"
	"os"
)

// 演示字符串读取
func TestReader(t *testing.T) {

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

// 演示文本替换
func TestReplacer(t *testing.T) {

	s := " Go Language "
	r := strings.NewReplacer("a", "A", "G", "g")
	fmt.Println(r.Replace(s)) // go LAnguAge

	r.WriteString(os.Stdout, s) // go LAnguAge

}


// 常用函数列表
func TestString(t *testing.T) {

	s := "this is a mock test demo, only test."

	// 子串
	fmt.Println(strings.EqualFold(s, "this is a mock test demo, only test.")) // true
	fmt.Println(strings.EqualFold(s, "This is a mock test demo, only test.")) // true

	fmt.Println(strings.HasPrefix(s, "this")) // true
	fmt.Println(strings.HasPrefix(s, "This")) // false

	fmt.Println(strings.Contains(s, "test")) // true
	fmt.Println(strings.Contains(s, "test001")) // false
	fmt.Println(strings.ContainsAny(s, "ta")) // true
	fmt.Println(strings.ContainsAny(s, "xz")) // false

	fmt.Println(strings.Count(s, "test")) // 2

	// 获取子串位置
	fmt.Println(strings.Index(s, "test")) // 15
	fmt.Println(strings.LastIndex(s, "test")) // 31
	fmt.Println(strings.Index(s, "test001")) // -1
	fmt.Println(strings.IndexByte(s, 's')) // 3
	fmt.Println(strings.IndexFunc(s, func(r rune) bool {
		return r == rune(97)
	})) // 8 对应字符a

	// 字符串中字符处理
	fmt.Println(strings.Title(s)) // This Is A Mock Test Demo, Only Test.
	fmt.Println(strings.ToLower(s)) // this is a mock test demo, only test.
	fmt.Println(strings.ToUpper(s)) // THIS IS A MOCK TEST DEMO, ONLY TEST.
	fmt.Println(strings.Repeat("test", 3)) // testtesttest
	fmt.Println(strings.Replace(s, "test", "TEST", -1)) // this is a mock TEST demo, only TEST.
	// 所有的字符编码+1
	fmt.Println(strings.Map(
		func(r rune) rune {
			return r + 1
		}, s)) // uijt!jt!b!npdl!uftu!efnp-!pomz!uftu/


	// 字符串前后端处理
	fmt.Println(strings.Trim(s, "test.")) // 感受一下 ^his is a mock test demo, only$
	fmt.Println(strings.TrimSpace(s)) // this is a mock test demo, only test.
	fmt.Println(strings.TrimFunc(s, func(r rune) bool {
		return r == 97
	})) // 只对首尾有效

	// 字符串分割与合并
	fmt.Println(strings.Fields(s)) // [this is a mock test demo, only test.]
	fmt.Println(strings.Split(s, "test")) // [this is a mock   demo, only  .]
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "$$$")) //a$$$b$$$c

}