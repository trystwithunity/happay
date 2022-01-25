package main

import (
	"github.com/trystwithunity/happay/queue"
)

func main() {
	go queue.EnqueueService(":5000")
	go queue.EnqueueService(":5001")
	go queue.EnqueueService(":5002")
	go queue.EnqueueService(":5003")
	go queue.EnqueueService(":5004")

	go queue.DequeueService()
	go queue.DequeueService()
	go queue.DequeueService()

	done := make(chan struct{})
	done <- struct{}{}
}
