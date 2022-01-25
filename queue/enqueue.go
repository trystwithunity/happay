package queue

import (
	"io"
	"log"
	"net/http"
	"time"
)

func EnqueueService(port string) {
	setupSink(port)
}

func setupSink(port string) {
	server := &http.Server{
		Addr:           port,
		Handler:        http.HandlerFunc(producerHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatalln(server.ListenAndServe())
}

// producerHandler handles producer requests and enqueues each request
// body of type unmarshalled json to the queue
func producerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Wrong request type, dropping request")
		return
	}
	json, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Could not read from request body, dropping request")
		return
	}
	defer r.Body.Close()
	// validation of json ommited for now
	getQueue().enqueue(json)
}
