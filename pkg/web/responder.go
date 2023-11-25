package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/htmx"
)

func Respond(w http.ResponseWriter, r *http.Request, item interface{}, t *template.Template) {
	if htmx.IsHTMX(r) {
		err := t.Execute(w, item)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	SendJson(w, r, item)
}

func RespondWithTemplate(w http.ResponseWriter, r *http.Request, item interface{}, t *template.Template, tmpl string) {
	if htmx.IsHTMX(r) {
		err := t.ExecuteTemplate(w, tmpl, item)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	SendJson(w, r, item)
}
