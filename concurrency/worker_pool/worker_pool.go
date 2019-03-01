package worker_pool

import "fmt"

type Task interface {
	Run()
}

type TaskExecutor struct {
	taskChan    chan Task
	numOfWorker int
}

func NewTaskExecutor(numOfWorker int) *TaskExecutor {
	taskChan := make(chan Task)
	for i := 0; i < numOfWorker; i++ {
		go func(id int) {
			for task := range taskChan {
				fmt.Printf("Worker %d is working.\n", id)
				task.Run()
			}
		}(i)
	}

	return &TaskExecutor{taskChan, numOfWorker}
}

func (ex *TaskExecutor) AddTask(task Task) {
	ex.taskChan <- task
}
