package main

import (
	"fmt"
	"os"
	"tasks/options"
)

func main() {
	option := options.Parse(os.Args[1:])
	fmt.Printf("%+v\n", option)
}
