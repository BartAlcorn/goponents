package minsse

import (
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
	index "github.com/bartalcorn/goponents/pkg/index"
	"github.com/bartalcorn/goponents/pkg/state"
	"github.com/bartalcorn/goponents/pkg/web"
)

// HandleMain handler the '/min/' GET route
func HandleMain(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "min"

	if htmx.IsHTMX(r) {
		web.Respond(w, r, nil, "pkg/minSse/tmpls/*.gohtml", "main")

	} else {
		index.Index(w, r)
	}

}
