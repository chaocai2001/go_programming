package concurrency

import "testing"

func TestCloseChannel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//		ch <- 10
	}()

	// for {
	// 	if data, ok := <-ch; ok {
	// 		t.Log(data)
	// 	} else {
	// 		break
	// 	}
	// }
	for data := range ch {
		t.Log(data)
	}
}
