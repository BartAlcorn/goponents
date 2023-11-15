package acquire

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"

	web "github.com/bartalcorn/goponents/pkg/web"
)

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	i, err := Find(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("error getting all", err)
	}

	t, err := template.ParseFiles("pkg/acquire/tmpls/details.gohtml", "pkg/acquire/tmpls/statusicons.gohtml")
	web.Responder(w, r, i, t)

}
