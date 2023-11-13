package web

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/web/state"
)

func ServerSentEvents(w http.ResponseWriter, r *http.Request) {

	state := web.State

	t, err := template.ParseFiles("web/index/sse.gohtml")
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
