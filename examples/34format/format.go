package main

import (
	"fmt"
	"os"
)

// 演示字符串格式化

// %+v   打印结构体时，会添加字段名     Printf("%+v", people)  {Name:zhangsan}
// %#v   相应值的Go语法表示            Printf("#v", people)   main.Human{Name:"zhangsan"}
// %T    相应值的类型的Go语法表示       Printf("%T", people)   main.Human
// %t    单词true或false
// %b    表示为二进制
// %c    该值对应的unicode码值
// %d    表示为十进制
// %o    表示为八进制
// %q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
// %x    表示为十六进制，使用a-f
// %X    表示为十六进制，使用A-F
// %U    表示为Unicode格式：U+1234，等价于"U+%04X"
// %b    无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat %e    科学计数法，如-1234.456e+78 %E    科学计数法，如-1234.456E+78 %f    有小数部分但无指数部分，如123.456 %F    等价于%f %g    根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
// %G    根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
// %s    直接输出字符串或者[]byte %q    该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
// %x    每个字节用两字符十六进制数表示（使用a-f）
// %X    每个字节用两字符十六进制数表示（使用A-F）
// %p    指针，表示为十六进制，并加上前导的0x
// %f:    默认宽度，默认精度
// %9f    宽度9，默认精度
// %.2f   默认宽度，精度2 %9.2f  宽度9，精度2 %9.f   宽度9，精度0

// 用于演示的结构体
type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)

	fmt.Printf("%+v\n", p)

	fmt.Printf("%#v\n", p)

	fmt.Printf("%T\n", p)

	fmt.Printf("%t\n", true)

	fmt.Printf("%d\n", 123)

	fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	fmt.Printf("%x\n", 456)

	fmt.Printf("%f\n", 78.9)

	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// 返回格式化后的字符串
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 将格式化后的字符串重定向到指定的流
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}


