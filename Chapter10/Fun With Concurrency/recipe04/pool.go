package main

import "sync"
import "fmt"
import "time"

type Worker struct {
	id string
}

func (w *Worker) String() string {
	return w.id
}

var globalCounter = 0

var pool = sync.Pool{
	New: func() interface{} {
		res := &Worker{fmt.Sprintf("%d", globalCounter)}
		globalCounter++
		return res
	},
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			// This code block is done only once
			w := pool.Get().(*Worker)
			fmt.Println("Got worker ID: " + w.String())
			time.Sleep(time.Second)
			pool.Put(w)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
