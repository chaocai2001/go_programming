package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() chan string {
	retCh := make(chan string)
	//retCh := make(chan string, 1)
	go func() {
		time.Sleep(time.Millisecond * 1) //Client will be stuck
		fmt.Println("Do something")
		retCh <- "result"
		fmt.Println("Done")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := service()
	fmt.Println("Do something else.")
	time.Sleep(time.Millisecond * 50) //service will be stuck
	fmt.Println("After some time consuming operation")
	ret := <-retCh
	t.Log(ret)
}
