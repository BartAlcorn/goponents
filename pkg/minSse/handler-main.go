package minSse

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Masterminds/sprig/v3"
)

// ShowMain handler the '/min/' GET route
func ShowMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("simulator.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
