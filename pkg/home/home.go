package home

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/pkg/state"
)

// Home serves the GET "/home/" route
func Home(w http.ResponseWriter, r *http.Request) {
	web.State.Module = "home"

	t, err := template.ParseGlob("pkg/home/*.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.ExecuteTemplate(w, "tech", &web.State)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
