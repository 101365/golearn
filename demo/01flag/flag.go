package main

import (
	"flag"
	"fmt"
)

var version string

func init() {
	flag.StringVar(&version, "version", "def_version", "")
}

// go run flag.go -version=123
func main() {
	flag.Parse()
	fmt.Println("version => ", version)
}
