package application

import (
	"net/http"
)

func SSESimulator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// names := []string{"Mike", "Bart", "Ben", "Jordan", "David"}

	// // Create a channel to send data
	// dataCh := make(chan string)

	// // Create a context for handling client disconnection
	// _, cancel := context.WithCancel(r.Context())
	// defer cancel()

	// // Send data to the client
	// go func() {
	// 	for data := range dataCh {
	// 		idx := rand.Intn(len(names))
	// 		fmt.Fprintf(w, "data: %s by %s\n\n", data, names[idx])
	// 		w.(http.Flusher).Flush()
	// 	}
	// }()

	// // Simulate sending data periodically
	// for {
	// 	dataCh <- time.Now().Format(time.TimeOnly)
	// 	time.Sleep(2 * time.Second)
	// }

}
