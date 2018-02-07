package schema

import (
	"regexp"
	"testing"
)

func TestContextFreeRegexp(t *testing.T) {
	for _, token := range ContextFreeTokenTypes {
		_, err := regexp.Compile(token.regexp)

		if err != nil {
			t.Errorf("Regex '%v' didn't compile for context free token %v", token.regexp, token.Name())
		}
		t.Failed()
	}
}
