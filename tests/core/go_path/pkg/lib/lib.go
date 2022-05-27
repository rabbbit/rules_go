package lib

import (
	"C"
	_ "embed" // for go:embed
)

//go:embed embedded_src.txt
var embeddedSource string

//go:embed template/index.html.tmpl
var indexTmpl string
