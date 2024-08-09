package dev

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/m3rashid/gotth-components/components"
	"github.com/m3rashid/gotth-components/components/page"
)

func StartServer() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	componentSpecs := components.ConfigWithComponentSpecs

	for componentName := range componentSpecs {
		component := components.ConfigWithComponentSpecs[componentName].Spec()
		handler := &templ.ComponentHandler{
			Component: page.Page(page.PageProps{Title: componentName, Children: component}),
		}
		http.Handle("/"+componentName, handler)
	}

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
