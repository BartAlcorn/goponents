package ssesimulator

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/gommon/color"
)

func Simulate(w http.ResponseWriter, r *http.Request) {
	color.Println(color.Yellow("Simulate Handler: started"))
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// w.Header().Set("Transfer-Encoding", "chunked")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	updateCh := make(chan Asset)

	// Create a context for handling client disconnection
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Send data to the client
	go processAssets(ctx, flusher, w)

	time.Sleep(1 * time.Second)
	// generateUpdates
	go generateUpdates(r.Context(), updateCh)
	fmt.Println("starting updates")

	for updateEvent := range updateCh {
		event, err := formatUpdate("min-event-update", updateEvent)
		if err != nil {
			fmt.Println(err)
			break
		}

		_, err = fmt.Fprint(w, event)
		if err != nil {
			fmt.Println(err)
			break
		}

		flusher.Flush()
	}

	color.Println(color.Red("Simulate Handler: ended"))

}
