package minSse

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/Masterminds/sprig/v3"
)

func formatDetails(event string, data any) (string, error) {
	var tpl bytes.Buffer
	t, err := template.New("simulator.gohtml").Funcs(sprig.FuncMap()).ParseGlob("pkg/minSse/tmpls/*.gohtml")
	if err != nil {
		fmt.Println("error parsin template formatAsset", err)
	}

	err = t.ExecuteTemplate(&tpl, "details", data)
	if err != nil {
		fmt.Println("error executing", err)
		// w.WriteHeader(http.StatusInternalServerError)
	}

	sb := strings.Builder{}
	sanitized := strings.ReplaceAll(tpl.String(), "\n  ", "")
	sanitized = strings.ReplaceAll(sanitized, "\n", "")
	sb.WriteString(fmt.Sprintf("event: %s\n", event))
	sb.WriteString(fmt.Sprintf("data: %v\n\n", sanitized))

	return sb.String(), nil
}
