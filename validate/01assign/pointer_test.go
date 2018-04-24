package main

import (
	"testing"
	"fmt"
)

func TestPointer(t *testing.T) {
	//
	//v := "hello"
	//
	//demo(&v)
	//
	//var x Demo = TwoPoint{}
	//fmt.Println(x)

	errorMethod()

	fmt.Println("after panic")

}

func errorMethod() string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	real()

	return "success"
}

//go:nosplit
func real() {
	panic("错误啦")
}

func demo(s *string)  {
	fmt.Println(s)
}

type Point struct {
	X int
}

type Point2 struct {
	X int
}

type TwoPoint struct {
	Demo
	Point
	//Point2
}

type Demo interface {

}
