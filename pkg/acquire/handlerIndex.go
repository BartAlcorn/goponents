package acquire

import (
	"fmt"
	"html/template"
	"net/http"
)

// Read calls the GetAll func
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/acquire/tmpls/index.gohtml")
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
