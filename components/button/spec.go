package button

import "github.com/a-h/templ"

type ButtonComponentSpec struct{}

func (b ButtonComponentSpec) Spec() templ.Component {
	// variants := ButtonVariantTypes
	// sizes := ButtonSizeTypes

	// buttons := make([]func() templ.Component, len(variants)*len(sizes))
	return Button(ButtonProps{Label: "Button", Variant: "primary"})

	// for _, variant := range variants {
	// 	for _, size := range sizes {
	// 		buttons = append(buttons, func() templ.Component {
	// 			return Button(ButtonProps{Size: size, Variant: variant})
	// 		})
	// 	}
	// }

	// return buttons
}
