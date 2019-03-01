package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type SingletonObj struct {
}

var once sync.Once
var obj *SingletonObj

func GetSingletonObj() *SingletonObj {
	once.Do(func() {
		fmt.Println("Create Singleton obj.")
		obj = &SingletonObj{}
	})
	return obj
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("Obj address: %d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
