package main

import (
	"os"
	"os/exec"

	"github.com/paganotoni/tailo"
)

func main() {
	tailo.Build(
		tailo.UseInputPath("./web/index.css"),
		tailo.UseConfigPath("./tailwind.config.js"),
		tailo.UseOutputPath("./static/goponents.css"),
	)

	cmd := exec.Command("go", "build")
	cmd.Args = append(
		cmd.Args,

		`--ldflags`, `-linkmode=external -extldflags="-static"`,
		`-tags`, `osusergo,netgo,musl`,
		`-buildvcs=false`,
		"-o", "bin/app",
		"./main.go",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
