package application

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func SSESimulator(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SSESimulator started")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	names := []string{"Mike", "Bart", "Ben", "Jordan", "David"}

	// Create a channel to send data
	minEventCh := make(chan string)

	// Create a context for handling client disconnection
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Send data to the client
	go func() {
		for data := range minEventCh {
			idx := rand.Intn(len(names))
			fmt.Fprintf(w, "data: %s by %s\n\n", data, names[idx])
			w.(http.Flusher).Flush()
		}
	}()

	// Simulate sending data periodically
	for {
		minEventCh <- time.Now().Format(time.TimeOnly)
		time.Sleep(2 * time.Second)
	}

}
