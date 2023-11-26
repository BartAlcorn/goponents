package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	"github.com/bartalcorn/goponents/pkg/htmx"
	web "github.com/bartalcorn/goponents/pkg/index"
)

// Respond takes in the .gohtml template file(s), parses the template(s)
// and return the result to the client:
// if content-type is "application/json", it returns json
// if HX-Request is "true", it returns the HTMX fragment
// else it returns the index page, which the global module var will set the correct "main"
func Respond(w http.ResponseWriter, r *http.Request, item interface{}, t string, block string) {
	var err error

	if r.Header.Get("content-type") == "application/json" {
		SendJson(w, r, item)
		return
	}

	tmpl, err := template.New("main.gohtml").Funcs(sprig.FuncMap()).ParseGlob(t)
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if htmx.IsHTMX(r) {
		if block == "" {
			err = tmpl.Execute(w, item)
		} else {
			err = tmpl.ExecuteTemplate(w, block, item)

		}
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	web.Index(w, r)

}
