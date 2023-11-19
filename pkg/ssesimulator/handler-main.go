package ssesimulator

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/web"
)

// Create calls the Insert func
func ShowMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/ssesimulator/tmpls/simulator.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	web.Respond(w, r, nil, t)
}
