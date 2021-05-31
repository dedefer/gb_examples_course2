package main

import (
	"sync"
)

var _ WP = &WorkerPool{}

type Task = func()

type TaskStruct struct {
	task Task
	done chan struct{}
}

type WorkerPool struct {
	tasksCh chan TaskStruct

	closeMu sync.RWMutex
	closed  bool
	wg      sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	tasksCh := make(chan TaskStruct, workers)

	wp := &WorkerPool{
		tasksCh: tasksCh,
	}

	wp.wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			defer wp.wg.Done()

			for task := range tasksCh {
				task.task()
				task.done <- struct{}{}
			}
		}(i)
	}

	return wp
}

func (wp *WorkerPool) Do(task Task) {
	wp.closeMu.RLock()
	defer wp.closeMu.RUnlock()
	if wp.closed {
		return
	}

	wp.tasksCh <- TaskStruct{
		task: task,
		done: make(chan struct{}, 1),
	}
}

func (wp *WorkerPool) DoBatch(tasks ...Task) {
	wp.closeMu.RLock()
	defer wp.closeMu.RUnlock()
	if wp.closed {
		return
	}

	doneCh := make(chan struct{})
	addedTasks := 0

	for _, task := range tasks {
		for taskAdded := false; !taskAdded; {
			select {
			case wp.tasksCh <- TaskStruct{
				task: task,
				done: doneCh,
			}:
				addedTasks++
				taskAdded = true
			case <-doneCh:
				addedTasks--
			}
		}
	}

	for ; addedTasks > 0; addedTasks-- {
		<-doneCh
	}
}

func (wp *WorkerPool) Close() {
	wp.closeMu.Lock()
	wp.closed = true
	wp.closeMu.Unlock()

	close(wp.tasksCh)

	wp.wg.Wait()
}
