package main

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	bytes := []byte("hello world")
	copy(bytes,"ha ha")
	fmt.Println(string(bytes))

	var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1, 100: 1}
	fmt.Println(asciiSpace)
	fmt.Println(int('\t'))
}
