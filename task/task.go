package task

import (
	"fmt"
	"math/rand"
	"time"
)

type TaskHandle[T any] interface {
	Random() int
	Detail() string
	Finished(int) // remove the finished task, then rewrite the task config file
}

type Task[T any] struct {
	Describe string
	Tasks    []T
	Cycle    bool
}

type DailyTask[T any] struct {
	Task[T]
	TaskType string
}

func NewDailyTask[T any](taskType string, tasks Task[T]) *DailyTask[T] {
	// TODO: remove finished task
	// if cycle is true and all tasks were finished, 
	// restart it from the beginning.
	return &DailyTask[T]{
		Task:     tasks,
		TaskType: taskType,
	}
}

func (d *DailyTask[T]) Detail(index int) string {
	return fmt.Sprintf(d.Describe, d.Tasks[index])
}

func (d *DailyTask[T]) Random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(len(d.Tasks))
}
