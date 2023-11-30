package minsse

import (
	"context"
	"fmt"
	"net/http"
)

func processAssets(ctx context.Context, flusher http.Flusher, w http.ResponseWriter) {

	// Send data to the client
	go generateAssets(8)
	// Send event data to the client
	for addEvent := range addCh {
		addEvent.Counts = Stats["counts"]
		addEvent.Metrics = Stats["metrics"]
		event, err := formatReturn("min-event-asset", addEvent, "assets")
		if err != nil {
			fmt.Println("ERROR: generateAssets: formatReturn", err)
			break
		}

		_, err = fmt.Fprint(w, event)
		if err != nil {
			fmt.Println("ERROR: generateAssets: Fprint", err)
			fmt.Println("ERROR: generateAssets: w", w)
			break
		}

		flusher.Flush()
	}

}
