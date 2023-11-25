package minsse

import (
	"context"
	"fmt"
	"net/http"
)

func processAssets(ctx context.Context, flusher http.Flusher, w http.ResponseWriter) {
	addCh := make(chan Asset)

	// Send data to the client
	go generateAssets(addCh)
	// Send event data to the client
	for addEvent := range addCh {
		event, err := formatReturn("min-event-asset", addEvent, "assets")
		if err != nil {
			fmt.Println("ERROR: generateAssets: formatReturn", err)
			break
		}

		_, err = fmt.Fprint(w, event)
		if err != nil {
			fmt.Println("ERROR: generateAssets: Fprint", err)
			break
		}

		flusher.Flush()
	}

}
