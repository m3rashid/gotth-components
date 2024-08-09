package utils

import (
	"github.com/a-h/templ"
)

type ComponentSpec interface {
	Spec() templ.Component
}
