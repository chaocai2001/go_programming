package panic_recover

import (
	"fmt"
	"testing"
)

//
// func TestPanic(t *testing.T) {
// 	defer func() {
// 		t.Log("program crashed")
// 	}()
// 	arr := [3]int{}
// 	for i := 0; i < 5; i++ {
// 		// if i == 1 {
// 		// 	os.Exit(-1)
// 		// }
// 		t.Log(arr[i])
// 	}
// }

func TestPanicRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover error:", err)
		}
	}()
	arr := [3]int{}
	for i := 0; i < 5; i++ {
		// if i == 1 {
		// 	os.Exit(-1)
		// }
		t.Log(arr[i])
	}

}
