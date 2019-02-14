package main

// import "fmt"
// import demo "github.com/chaocai2001/go_programming/demos"

import (
	"fmt"

	"github.com/chaocai2001/go_programming/demos"
)

func main() {
	n := 5
	// fmt.Println(os.Args)
	// if len(os.Args) > 1 {
	// 	if i, err := strconv.Atoi(os.Args[1]); err == nil {
	// 		n = i
	// 	}
	// }
	fmt.Println(demos.GetFibonacci(n))
}
