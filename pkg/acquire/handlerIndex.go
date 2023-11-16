package acquire

import (
	"fmt"
	"html/template"
	"net/http"

	web "github.com/bartalcorn/goponents/pkg/web"
)

// Read calls the GetAll func
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/acquire/tmpls/index.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	web.Respond(w, r, nil, t)
}
