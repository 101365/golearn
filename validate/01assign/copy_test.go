package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	bytes := []byte("hello world")
	copy(bytes,"ha ha")
	fmt.Println(string(bytes))
}
