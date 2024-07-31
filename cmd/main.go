package main

import (
	"flag"

	"github.com/m3rashid/gotth-components/dev"
	"github.com/m3rashid/gotth-components/docs"
)

func main() {
	genJSON := flag.Bool("gen-json", false, "Generate JSON file")
	startDocsServer := flag.Bool("docs", false, "Start the docs server")
	// ... more flags

	flag.Parse()

	if *genJSON {
		generateJsonConfig()
	} else if *startDocsServer {
		docs.StartServer()
	} else {
		dev.StartServer()
	}
}
