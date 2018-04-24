package main

import (
	"testing"
	"fmt"
)

func TestPointer(t *testing.T) {

	v := "hello"

	demo(&v)

}

func demo(s *string)  {
	fmt.Println(s)
}
