# Assignment

### package producer produces 100 requests for mutiple queuing services running at ports [5000:5004]

### package queue contains the core queuing service, it exposes two functions
- time stamps are time.Duration values since start of queuing service, can be replaced by absolute timestamps
- configurable wait time has been left out (consumers wait for a second when they have nothing to read)

### package consumer contains an endpoint to receive info from dequeuing services at port 8000

### unit tests are omitted due to lack of time, the modules have been checked using "go race -run"

- Start consumer "go run -race consumer.go"
- Start Queuing service "go run -race main.go"
- Start producer "go run -race producer.go"

