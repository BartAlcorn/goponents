package ssesimulator

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
		event, err := formatUpdate("min-event-asset", addEvent)
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
