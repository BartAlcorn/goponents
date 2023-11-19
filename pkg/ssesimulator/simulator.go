package ssesimulator

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/gommon/color"
)

func Simulate(w http.ResponseWriter, r *http.Request) {
	color.Println(color.Red("Simulate STARTED"))
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// w.Header().Set("Transfer-Encoding", "chunked")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	minEventCh := make(chan Asset)

	// Create a context for handling client disconnection
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Send data to the client
	go generateUpdate(r.Context(), minEventCh)

	// Send event data to the client
	for minevent := range minEventCh {
		Assets = append(Assets, minevent)
		event, err := formatServerSentEvent("min-event-update", minevent)
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

}
