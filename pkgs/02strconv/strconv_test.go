package _2strconv

import (
	"testing"
	"strconv"
	"fmt"
)

// bool 类型转换
func TestBool(t *testing.T)  {
	s := "true"
	b, _ := strconv.ParseBool(s)
	fmt.Println(b)
}

func TestInit(t *testing.T)  {
	fmt.Println(strconv.ParseInt("FF", 16, 0))
	// 255

	// 8进制字符串中只能包含0～7
	fmt.Println(strconv.ParseInt("FF", 8, 0))
	// 0 strconv.ParseInt: parsing "FF": invalid syntax

	// 下面错误说明：当前缀中包含0x这种字符时，必须指定base为0才不会报错。
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
	// 255
	fmt.Println(strconv.ParseInt("9", 10, 4))
	// 7 strconv.ParseInt: parsing "9": value out of range
}

func TestFloat(t *testing.T)  {

	// 转换结果会有精度问题，但是符合 IEEE754标准
	s := "0.12345678901234567890"

	f, err := strconv.ParseFloat(s, 32)
	fmt.Println(f, err)                // 0.12345679104328156
	fmt.Println(float32(f), err)       // 0.12345679

	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err)				   // 0.12345678901234568

}

func TestAppend(t *testing.T)  {

	dst := make([]byte, 0)
	// 将数字12按照二进制转换为字节，添加到dst中
	dst = strconv.AppendInt(dst, 12, 2)
	fmt.Println(dst) // [49 49 48 48]
	fmt.Println(string(dst)) // 1100

	// 在ascii码中，49对应数字1，48对应数字0
	// 二进制1100转换为10进制就是12

}
