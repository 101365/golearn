package main

import (
	"testing"
	"unsafe"
	"fmt"
)

func TestUnsafePointer(t *testing.T) {

	s := struct {
		a byte
		b byte
		c byte
		d int64
	}{0, 0, 0, 0}

	// 将结构体指针转换为通用指针
	p := unsafe.Pointer(&s)
	// 保存结构体的地址备用（偏移量为 0）
	up0 := uintptr(p)
	// 将通用指针转换为 byte 型指针
	pb := (*byte)(p)
	// 给转换后的指针赋值
	*pb = 10
	// 结构体内容跟着改变
	fmt.Println(s)

	// 偏移到第 2 个字段
	up := up0 + unsafe.Offsetof(s.b)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 20
	// 结构体内容跟着改变
	fmt.Println(s)

	fmt.Println(pb)
	fmt.Println(&s.b)

	// 偏移到第 3 个字段
	up = up0 + unsafe.Offsetof(s.c)
	fmt.Println(unsafe.Alignof(s.c))
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 byte 型指针
	pb = (*byte)(p)
	// 给转换后的指针赋值
	*pb = 30
	// 结构体内容跟着改变
	fmt.Println(s)


	fmt.Println(pb)
	fmt.Println(&s.c)

	// 偏移到第 4 个字段
	up = up0 + unsafe.Offsetof(s.d)
	// 将偏移后的地址转换为通用指针
	p = unsafe.Pointer(up)
	// 将通用指针转换为 int64 型指针
	pi := (*int64)(p)
	// 给转换后的指针赋值
	*pi = 40
	// 结构体内容跟着改变
	fmt.Println(s)
	fmt.Println(unsafe.Alignof(s.d))
	fmt.Println(p)
	fmt.Println(&s.d)

}

func TestAlign(t *testing.T) {

	fmt.Println("|", "类型", "|", "对齐值")
	fmt.Println("|", "byte", "|", unsafe.Alignof(byte(0)))
	fmt.Println("|", "int8", "|", unsafe.Alignof(int8(0)))
	fmt.Println("|", "unit8", "|", unsafe.Alignof(uint8(0)))
	fmt.Println("|", "int16", "|", unsafe.Alignof(int16(0)))
	fmt.Println("|", "uint16", "|", unsafe.Alignof(uint16(0)))
	fmt.Println("|", "int32", "|", unsafe.Alignof(int32(0)))
	fmt.Println("|", "uint32", "|", unsafe.Alignof(uint32(0)))
	fmt.Println("|", "int64", "|", unsafe.Alignof(int64(0)))
	fmt.Println("|", "uint64", "|", unsafe.Alignof(uint64(0)))
	fmt.Println("|", "unitptr", "|", unsafe.Alignof(uintptr(0)))
	fmt.Println("|", "float32", "|", unsafe.Alignof(float32(0)))
	fmt.Println("|", "float64", "|", unsafe.Alignof(float64(0)))
	fmt.Println("|", "complex64", "|", unsafe.Alignof(complex64(0)))
	fmt.Println("|", "complex64", "|", unsafe.Alignof(complex64(0)))
	fmt.Println("|", "string", "|", unsafe.Alignof(""))
	fmt.Println("|", "int", "|", unsafe.Alignof(new(int)))
	fmt.Println("|", "struct", "|", unsafe.Alignof(struct {
		f  float32
		ff float64
	}{}))
	fmt.Println("|", "chan", "|", unsafe.Alignof(make(chan bool, 10)))
	fmt.Println("|", "[]int", "|", unsafe.Alignof(make([]int, 10)))
	fmt.Println("|", "map", "|", unsafe.Alignof(make(map[string]string, 10)))

}
