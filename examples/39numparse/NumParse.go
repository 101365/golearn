package main

import (
	"strconv"
	"fmt"
)

// 演示数字解析
// 在go语言中，字符串解析为数字，主要使用strconv包来完成
func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 第二个参数0代表自动推断字符串所代表的数字的进制，第三个参数64说明返回的整形数是以64位存储的
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// 返回无符号整数
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
