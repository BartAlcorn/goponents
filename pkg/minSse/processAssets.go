package minsse

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
	"github.com/kylelemons/godebug/pretty"
)

func processAssets(ctx context.Context, flusher http.Flusher, w http.ResponseWriter) {

	// Send data to the client
	go generateAssets(8)
	// Send event data to the client
	for addEvent := range addCh {
		addEvent.Counts = Stats["counts"]
		addEvent.Metrics = Stats["metrics"]
		event, err := htmx.SSERespond(addEvent, "pkg/minsse/tmpls/*.gohtml", "assets", "min-event-asset")
		if err != nil {
			fmt.Println("ERROR: generateAssets: SSERespond", err)
			break
		}

		_, err = fmt.Fprint(w, event)
		if err != nil {
			fmt.Println("ERROR: generateAssets: Fprint", err)
			fmt.Println("ERROR: generateAssets: w")
			pretty.Print(w)
			break
		}

		flusher.Flush()
	}

}
