package schema

import "testing"

func TestSemanticModelConsistency(t *testing.T) {
	for _, ttype := range SemanticTokenTypes {
		if ttype.PosModel().Binding == BindRight {
			sibling := ttype.Variant().OptValueTokenType()
			if sibling == nil {
				t.Errorf("%v : a bound-right token type must have an opt value token type associated with its variant.", ttype.Name())
			}
			if !ttype.PosModel().IsOptionFlag {
				t.Errorf("%v : a bound-right token type should be an option flag. ", ttype.Name())
			}
		}
		if ttype.PosModel().Binding == BindLeft {
			sibling := ttype.Variant().FlagTokenType()
			if sibling == nil {
				t.Errorf("%v : a bound-left token type must have a flag token type associated with its variant.", ttype.Name())
			}
			if ttype.PosModel().IsOptionFlag {
				t.Errorf("%v : a bound-left token type cannot be an option flag.", ttype.Name())
			}
			if !ttype.PosModel().IsOptionPart {
				t.Errorf("%v : a bound-left token type should be an option flag.", ttype.Name())
			}
		}
		if ttype.PosModel().Binding == BindUnknown {
			t.Errorf("%v : a semantic token type cannot have an unknown binding.", ttype.Name())
		}
	}
}
