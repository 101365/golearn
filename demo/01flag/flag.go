package main

import (
	"flag"
	"fmt"
)

var version string

func init() {
	flag.StringVar(&version, "version", "def_version", "test")
}

// go run flag.go -version=123
// go run flag.go --help
func main() {
	flag.Parse()
	fmt.Println("version => ", version)
}
