package main

import (
	_ "embed"

	"github.com/polygomic/polygomic-edge/command/root"
	"github.com/polygomic/polygomic-edge/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
