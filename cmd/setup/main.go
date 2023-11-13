package main

import (
	"fmt"

	"github.com/paganotoni/tailo"
)

func main() {
	err := tailo.Setup()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("✅ Tailwind CSS setup successfully")
}
