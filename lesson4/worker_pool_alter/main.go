package main

import (
	"context"
	"fmt"
	"sort"
)

type Msg struct {
	id   int
	data int
}

type Res struct {
	id  int
	res string
}

func main() {

	workersCnt := 3

	msgCh := make(chan Msg, workersCnt)
	resCh := make(chan Res, workersCnt)

	for i := 0; i < workersCnt; i++ {
		go func(gouroutineID int) {
			for message := range msgCh {
				resCh <- Res{
					id:  message.id,
					res: fmt.Sprintf("goroutine %d: processed %+v", gouroutineID, message),
				}
			}
		}(i)
	}

	// заранее знаем сколько сообщений
	msgCnt := 30

	// запускаем горутину для чтения в массив
	resps := make([]Res, 0, msgCnt)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for len(resps) < msgCnt {
			resps = append(resps, <-resCh)
		}
	}()

	// посылаем сообщения
	for i := 0; i < msgCnt; i++ {
		msgCh <- Msg{id: i, data: i * 100}
	}
	close(msgCh) // больше не будем писать
	<-done       // ждем пока все прочитается

	// сортируем по id
	sort.Slice(resps, func(i, j int) bool {
		return resps[i].id < resps[j].id
	})

	// выводим на экран
	for _, res := range resps {
		fmt.Println(res.res)
	}

}

type Task = func()

type WorkerMsg struct {
	task    Task
	resChan chan struct{}
}

type BatchWorkerPool struct {
	workers int
	msgChan chan WorkerMsg
}

// Эта штука создает воркер пул с {workers} воркеров
func New(workers int) *BatchWorkerPool {
	msgChan := make(chan WorkerMsg, workers)
	for i := 0; i < workers; i++ {
		go func() {
			for msg := range msgChan {
				msg.task()
				msg.resChan <- struct{}{}
			}
		}()
	}
	return &BatchWorkerPool{
		workers: workers,
		msgChan: msgChan,
	}
}

// Эта функция принимает таски, исполняет их в пуле и ждет, пока все таски не завершатся
func (p *BatchWorkerPool) ExecuteBatch(tasks ...Task) {
	ctx, cancel := context.WithCancel(context.Background())
	resChan := make(chan struct{}, p.workers)
	go func() {
		defer cancel()
		for tasksCount := len(tasks); tasksCount > 0; tasksCount-- {
			<-resChan
		}
	}()

	for _, task := range tasks {
		p.msgChan <- WorkerMsg{
			resChan: resChan,
			task:    task,
		}
	}

	<-ctx.Done()
}
