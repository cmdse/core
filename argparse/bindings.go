package argparse

import (
	. "cmdse-cli/schema"
)

type Bindings []Binding

func (bindings Bindings) Contains(bindingToCheck Binding) bool {
	for _, binding := range bindings {
		if binding == bindingToCheck {
			return true
		}
	}
	return false
}
