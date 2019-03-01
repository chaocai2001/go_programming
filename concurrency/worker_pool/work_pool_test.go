package worker_pool

import (
	"fmt"
	"testing"
	"time"
)

type MyTask struct {
	TaskID int
}

func (t *MyTask) Run() {
	time.Sleep(time.Millisecond * 10)
	fmt.Printf("Task %d is done.\n", t.TaskID)
}

func TestWorkPool(t *testing.T) {
	exec := NewTaskExecutor(3)
	for i := 0; i < 10; i++ {
		exec.AddTask(&MyTask{i})
	}
}
