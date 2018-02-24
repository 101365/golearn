package main

import (
	"fmt"
	"reflect"
	"math"
)

// 演示常量的定义

const s  = "hello"

const (
	a = "a"
	b = 12
)

func main() {

	/*
		1.const用于声明一个常量
		2.const可以出现在任何var可以出现的地方
	 */

	const k  = "world"

	fmt.Println(s)
	fmt.Println(k)
	fmt.Println(a, b)

	const n  = 100000000

	// 常数表达式可以执行任意精度的运算
	const m  = 1e25 / n
	fmt.Println(m)

	//
	fmt.Println(reflect.TypeOf(n)) // int
	fmt.Println(reflect.TypeOf(m)) // float64

	fmt.Println(math.Sin(m))



}
