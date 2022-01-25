package queue

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func DequeueService() {
	q := getQueue()
	for {
		select {
		case item := <-q:
			// TODO: check if url is available
			_, err := http.Post("http://localhost:8000", "application/text", bytes.NewBuffer([]byte(item.String())))
			if err != nil {
				fmt.Printf("Error while posting to consumer %v\n", err)
			}
		default:
			// TODO: configure for custom time
			fmt.Println("Sleeping for 1 sec, nothing to read")
			time.Sleep(time.Second)
		}
	}
}
