package argparse

import (
	. "cmdse-cli/schema"
	"testing"
)

func TestBindings_Contains(t *testing.T) {
	bindings := Bindings{LEFT, RIGHT}
	if !bindings.Contains(LEFT) {
		t.Errorf("Bindings should contain 'LEFT'")
	}
	if !bindings.Contains(RIGHT) {
		t.Errorf("Bindings should contain 'RIGHT'")
	}
	if bindings.Contains(NONE) {
		t.Errorf("Bindings should not contain 'NONE'")
	}
}
