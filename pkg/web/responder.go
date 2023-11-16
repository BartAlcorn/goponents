package web

import (
	"fmt"
	"html/template"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, item interface{}, t *template.Template) {
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX == "true" {
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
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX == "true" {
		err := t.ExecuteTemplate(w, tmpl, item)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	SendJson(w, r, item)
}
