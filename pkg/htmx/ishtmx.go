package htmx

import "net/http"

// IsHTMX returns bool indicating the presence of the
// "HX-Request" Request Header with a value of "true"
// This is useful in determining what and how to return
func IsHTMX(r *http.Request) bool {

	return r.Header.Get("HX-Request") == "true"
}
