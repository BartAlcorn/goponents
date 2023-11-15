package acquire

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

// Read calls the GetAll func
func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/acquire/tmpls/acquireIndex.gohtml")
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

func ModalBtn(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseGlob("pkg/acquire/tmpls/acquireModalDetails.gohtml")
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

func StatusDetails(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pkg/acquire/tmpls/acquireStatusIcon.gohtml", "pkg/acquire/tmpls/statusicons.gohtml")
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

func BriefDetails(w http.ResponseWriter, r *http.Request) {
	item, err := Find(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("error getting all", err)
	}

	isHTMX := r.Header.Get("HX-Request")
	if isHTMX == "true" {
		t, err := template.ParseFiles("pkg/acquire/tmpls/acquireBriefDetails.gohtml", "pkg/acquire/tmpls/statusicons.gohtml", "pkg/acquire/tmpls/discologo.gohtml")
		if err != nil {
			fmt.Println("error parsing gohtml", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		err = t.Execute(w, item)
		if err != nil {
			fmt.Println("error executing", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	jsn := fmt.Sprintf(
		"{\"status\": \"%s\", \"code\": \"%s\", \"detail\": \"%s\", \"date\": \"%s\"}",
		item.Asset.Status.Type,
		item.Asset.Status.Codes[0].Code,
		item.Asset.Status.Codes[0].Detail,
		item.Audit.CreatedAt,
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, err = w.Write([]byte(jsn))
	if err != nil {
		fmt.Println("failed to write", err)
	}
	return
}

// Read calls the GetAll func
func Read(w http.ResponseWriter, r *http.Request) {
	item, err := Find(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println("error getting all", err)
	}

	// if not HTMX, return JSON
	isHTMX := r.Header.Get("HX-Request")
	if isHTMX != "true" {
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

	// else return HTMX
	t, err := template.ParseFiles("pkg/acquire/tmpls/acquireDetails.gohtml", "pkg/acquire/tmpls/statusicons.gohtml")
	if err != nil {
		fmt.Println("error parsing gohtml", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = t.Execute(w, item)
	if err != nil {
		fmt.Println("error executing", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
