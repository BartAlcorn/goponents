package minsse

import (
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
	"github.com/bartalcorn/goponents/pkg/index"
	"github.com/bartalcorn/goponents/pkg/state"
)

// HandleMain handler the '/min/' GET route
func HandleMain(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "min"

	if htmx.IsHTMX(r) {
		htmx.Respond(w, r, nil, "pkg/minSse/tmpls/*.gohtml", "main")

	} else {
		index.Index(w, r)
	}

}
