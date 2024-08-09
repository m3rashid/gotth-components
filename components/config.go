package components

import (
	"github.com/m3rashid/gotth-components/components/button"
	"github.com/m3rashid/gotth-components/utils"
)

// map of component names to their dependencies
var ConfigWithDependencies = map[string][]string{
	"button": {},
}

var ConfigWithComponentSpecs = map[string]utils.ComponentSpec{
	"button": button.ButtonComponentSpec{},
}
