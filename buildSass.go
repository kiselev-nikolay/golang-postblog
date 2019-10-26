package main

import (
	"os"

	sass "github.com/wellington/go-libsass"
)

func BuildSass(dist string) {
	inputSass, _ := os.Open("_data/sass/bulma.sass")
	output, _ := os.Create(dist)
	comp, _ := sass.New(output, inputSass)
	includePaths := []string{"./sass/"}
	_ = comp.Option(
		sass.IncludePaths(includePaths),
		sass.WithSyntax(sass.SassSyntax),
		sass.OutputStyle(sass.COMPRESSED_STYLE))
	_ = comp.Run()
}
