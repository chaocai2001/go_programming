package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hi,", os.Args[1])
	}
	fmt.Println("Hello World!")
	os.Exit(3)
}
