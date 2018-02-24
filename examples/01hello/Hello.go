package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, world!!!")
	fmt.Println("Current runtime go version: ", runtime.Version())
}
