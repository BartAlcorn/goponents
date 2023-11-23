package state

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetThemes reads the ./tailwind.config.js file and
// returns the list of install themes
func GetThemes() (results []string) {
	file, err := os.Open("./tailwind.config.js")
	if err != nil {
		fmt.Println("ERROR opening file", err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 7 && line[:7] == "themes:" {
			l := len(line) - 2

			themes := strings.Split(line[9:l], ", ")
			// results := []string{}
			for _, t := range themes {
				t := strings.Trim(t, "\"")
				results = append(results, t)
			}
			return results
		}
	}

	defer file.Close()
	return results
}
