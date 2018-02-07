package argparse

import (
	. "cmdse-cli/schema"
	"testing"
)

func TestBindings_Contains(t *testing.T) {
	bindings := Bindings{BindLeft, BindRight}
	if !bindings.Contains(BindLeft) {
		t.Errorf("Bindings should contain 'BindLeft'")
	}
	if !bindings.Contains(BindRight) {
		t.Errorf("Bindings should contain 'BindRight'")
	}
	if bindings.Contains(BindNone) {
		t.Errorf("Bindings should not contain 'BindNone'")
	}
}
