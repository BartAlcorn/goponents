package acquire

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bartalcorn/goponents/pkg/web"
)

func ModalBtn(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/acquire/tmpls/modalDetails.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	web.Respond(w, r, nil, t)
}
