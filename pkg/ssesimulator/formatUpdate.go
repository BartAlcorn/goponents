package ssesimulator

import (
	"bytes"
	"fmt"
	"strings"
)

func formatServerSentEvent(event string, data any) (string, error) {
	var tpl bytes.Buffer
	err := t.Execute(&tpl, data)
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
