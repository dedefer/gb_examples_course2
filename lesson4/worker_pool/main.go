package main

import (
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

	n := 3

	ch := make(chan Msg, n)
	resCh := make(chan Res, n)

	for i := 0; i < n; i++ {
		go func(gouroutineID int) {
			for message := range ch {
				resCh <- Res{
					id:  message.id,
					res: fmt.Sprintf("goroutine %d: processed %+v", gouroutineID, message),
				}
			}
		}(i)
	}

	// заранее знаем сколько сообщений
	msgCnt := 30

	// считаем сколько послали в канал
	sent := 0
	resps := make([]Res, 0, msgCnt)
	// посылаем и одновременно читаем до тех пор пока ВСЕ НЕ ПОШЛЕМ
	for sent < msgCnt {
		select {
		case ch <- Msg{id: sent, data: sent * 100}:
			sent++
		case res := <-resCh:
			resps = append(resps, res)
		}
	}
	close(ch)

	// дочитываем оставшиеся сообщения из канала результатов
	for len(resps) < msgCnt {
		res := <-resCh
		resps = append(resps, res)
	}

	// сортируем по id
	sort.Slice(resps, func(i, j int) bool {
		return resps[i].id < resps[j].id
	})

	for _, res := range resps {
		fmt.Println(res.res)
	}

}
