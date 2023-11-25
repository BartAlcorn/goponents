package minsse

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/bartalcorn/goponents/pkg/htmx"
	index "github.com/bartalcorn/goponents/pkg/index"
	"github.com/bartalcorn/goponents/pkg/state"
)

// HandleMain handler the '/min/' GET route
func HandleMain(w http.ResponseWriter, r *http.Request) {
	state.State.Module = "min"

	if htmx.IsHTMX(r) {
		t, err := template.New("simulator.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
		err = t.Execute(w, nil)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		index.Index(w, r)
	}

}
