package web

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/web/state"
)

func Index(w http.ResponseWriter, r *http.Request) {

	state := web.State

	t, err := template.ParseGlob("web/index/*.gohtml")
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
