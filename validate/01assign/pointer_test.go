package main

import (
	"testing"
	"fmt"
)

func TestPointer(t *testing.T) {

	v := "hello"

	demo(&v)

	var x Demo = TwoPoint{}
	fmt.Println(x)

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
