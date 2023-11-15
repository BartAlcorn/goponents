package acquire

import (
	"fmt"
	"html/template"
	"net/http"
)

func StatusDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/acquire/tmpls/statusIcon.gohtml", "pkg/acquire/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
