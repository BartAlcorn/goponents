package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func Responder(w http.ResponseWriter, r *http.Request, item interface{}, t *template.Template) {
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX == "true" {
		err := t.Execute(w, item)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res, err := json.Marshal(item)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write(res)
	if err != nil {
		fmt.Println("failed to write", err)
	}
	return
}
