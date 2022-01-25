package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 8000")
	log.Fatalln(http.ListenAndServe(":8000", nil))
}

var counter int32 = 0

// handler for receiving responses from consumers
// displays total number of responses on get request
// handles text sent by dequeuing services
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler invoked")
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading response from queueing service: %v\n", err)
			return
		}
		atomic.AddInt32(&counter, 1)
		fmt.Printf("%v\n", string(body))
	}
	if r.Method == http.MethodGet {
		w.Write([]byte(strconv.Itoa(int(counter))))
	}
}
