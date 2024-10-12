package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

// var mu sync.Mutex

type ConcQueue struct {
	queue []int32
	mu    sync.Mutex
}

func (q *ConcQueue) enQueue(ele int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, ele)
}

func (q *ConcQueue) deQueue() int32 {
	if len(q.queue) == 0 {
		panic("Queue is Empty")
	}
	del := q.queue[0]
	q.queue = q.queue[1:]
	return del
}

func (q *ConcQueue) size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.queue)
}

func main() {
	q := ConcQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			q.enQueue(int32(rand.Intn(math.MaxInt32)))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(q.size())

}
