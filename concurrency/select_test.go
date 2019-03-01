package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func slowService() chan string {
	retCh := make(chan string)
	//retCh := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Do something")
		retCh <- "result"
		fmt.Println("Done")
	}()
	return retCh
}

func TestAsynService1(t *testing.T) {
	retCh := slowService()
	fmt.Println("Do something else.")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("After some time consuming operation")
	select {
	case ret := <-retCh:
		t.Log(ret)
	case <-time.After(time.Second * 1):
		t.Log("Time out")
	}

}
