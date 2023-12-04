package htmx

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
	// "github.com/bartalcorn/goponents/pkg/index"
)

// Respond takes in the .gohtml template file(s), parses the template(s)
// and return the result to the client:
// if content-type is "application/json", it returns json
// if HX-Request is "true", it returns the HTMX fragment
// else it returns a full, stand alone page, if the block is "main"
func Respond(w http.ResponseWriter, r *http.Request, item interface{}, t string, block string) {
	var err error

	if r.Header.Get("content-type") == "application/json" {
		SendJson(w, r, item)
		return
	}

	// allows inclusion of files specific files from outside the given template path
	baseTemplate, err := template.ParseFiles(
		"pkg/index/tmpls/scripts.gohtml")
	if err != nil {
		panic(err)
	}

	// add spring to the template and then parse the templates in path
	tmpl, err := template.Must(baseTemplate.Clone()).Funcs(sprig.FuncMap()).ParseGlob(t)
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	if IsHTMX(r) {
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
	// not IsHTMXm but IS "main", serve it as a stand alone, full page.
	if block == "main" {
		err = tmpl.ExecuteTemplate(w, "full", item)
	}

}
