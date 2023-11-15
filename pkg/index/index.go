package web

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/pkg/webstate"
)

func Index(w http.ResponseWriter, r *http.Request) {

	state := web.State

	t, err := template.ParseGlob("pkg/index/*.gohtml")
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
