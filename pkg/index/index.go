package index

import (
	"fmt"
	"html/template"
	"net/http"

	state "github.com/bartalcorn/goponents/pkg/state"
)

func Index(w http.ResponseWriter, r *http.Request) {

	newState := state.State

	t, err := template.ParseGlob("pkg/index/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = t.Execute(w, newState)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
