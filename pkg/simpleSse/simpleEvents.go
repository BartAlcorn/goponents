package simplesse

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/bartalcorn/goponents/pkg/state"
	"github.com/labstack/gommon/color"
)

func SimpleSse(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "sse"
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	color.Println(color.Green("starting Simple SSE"))

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "SSE not supported", http.StatusInternalServerError)
		return
	}

	// Create a channel for data
	simulateEventCh := make(chan string)

	// Create a context for handling client disconnection
	_, cancel := context.WithCancel(r.Context())
	defer cancel()

	// go routine for loop
	go minSseLoop(r.Context(), simulateEventCh)

	// for each event received
	for simulateEvent := range simulateEventCh {
		_, err := fmt.Fprint(w, simulateEvent)
		if err != nil {
			fmt.Println(err)
			break
		}
		flusher.Flush()
	}

}

func minSseLoop(ctx context.Context, simulateEventCh chan<- string) {
	ticker := time.NewTicker(time.Second)
	names := []string{"Allie", "Bart", "Charlene", "Doug", "Ellen", "Frank", "Glory", "Henry", "Ida", "John", "Kathleen", "Liam", "Mary", "Ned", "Ophelia", "Paul", "Quincy", "Robert", "Sue", "Tom", "Uma", "Victor", "Wilma", "Xander", "Yvette", "Zod"}

eventLoop:
	for {
		select {
		case <-ctx.Done():
			break eventLoop
		case <-ticker.C:
			idx := rand.Intn(len(names))
			result := fmt.Sprintf("data: %s by %s\n\n", time.Now().Format(time.TimeOnly), names[idx])
			simulateEventCh <- result
		}
	}

	ticker.Stop()

	close(simulateEventCh)

	color.Println(color.Red("stopping Simple SSE"))
}
