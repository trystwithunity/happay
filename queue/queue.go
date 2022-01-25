package queue

import (
	"fmt"
	"sync"
	"time"
)

// type item is the type of the FIFO queue
// each item in the queue has a timestamp
// associated with it along with the json
type item struct {
	json      []byte
	timestamp time.Duration
}

// String pretty prints item type of the queue
func (i item) String() string {
	return fmt.Sprintf("%s \t %v\n", i.timestamp, string(i.json))
}

// Queue represents a FIFO queue for json
type queue chan item

var (
	q     queue
	lock  = sync.Once{}
	start time.Time
)

// getQueue returns singleton
func getQueue() queue {
	lock.Do(func() {
		fmt.Println("Initiliazing queue")
		q = make(queue, 1000)
		start = time.Now()
	})
	return q
}

func (q queue) enqueue(json []byte) {
	i := item{json: json}
	i.timestamp = time.Duration(time.Since(start).Nanoseconds())
	fmt.Printf("%v\n", i)
	q <- i
}
