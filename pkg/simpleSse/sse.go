package simplesse

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/state"
)

func SimpleSseExample(w http.ResponseWriter, r *http.Request) {

	state := state.State

	t, err := template.ParseGlob("pkg/simplesse/tmpls/sse.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, state)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
