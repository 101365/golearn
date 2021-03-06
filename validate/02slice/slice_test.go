package main

import (
	"testing"
	"fmt"
)

func TestSlice(t *testing.T) {

	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)

}

type path []byte

func (p *path) ToUpper() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func TestExtend(t *testing.T)  {
	var iBuffer [10]int
	slice := iBuffer[0:0]
	for i := 0; i < 20; i++ {
		slice = Extend(slice, i)
		fmt.Println(slice)
		if cap(slice) == len(slice) {
			fmt.Println("slice is full!")
			break
		}
	}

	slice = make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

func TestNil(t *testing.T)  {
	var a []int
	fmt.Println(a == nil)
	a = append(a, 1, 2, 3)
	fmt.Println(a)

	var b [32]int
	fmt.Println(len(b))
	b[3] = 123
	fmt.Println(b)
}

func TestPrint(t *testing.T)  {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("%+q \n", sample)
}

func TestUTF8(t *testing.T)  {
	// 变量名可以是中文
	var 中文变量 = "中文字符"
	fmt.Println(中文变量)

	var 中文变量2 = 12
	fmt.Println(中文变量2)

}
