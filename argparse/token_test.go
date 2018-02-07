package argparse

import (
	. "cmdse-cli/schema"
	"testing"
)

func TestToken_IsBoundToOneOfCF(t *testing.T) {
	// Test with a context-free type
	var tokens = TokenList{}
	var tokenType TokenType = CfOneDashWord
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: CfOneDashWord.SemanticCandidates,
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundNoneOrRight := token.IsBoundToOneOf(Bindings{BindNone, BindRight})
	if !boundNoneOrRight {
		t.Errorf("token should be bound BindNone or BindRight")
	}

	boundNoneOrLeft := token.IsBoundToOneOf(Bindings{BindNone, BindLeft})
	if boundNoneOrLeft {
		t.Errorf("token should not be bound BindNone or BindLeft")
	}
	boundUnknownOrLeft := token.IsBoundToOneOf(Bindings{BindUnknown, BindLeft})
	if boundUnknownOrLeft {
		t.Errorf("token should not be bound BindUnknown or BindLeft")
	}
}

func TestToken_IsBoundToOneOfSem(t *testing.T) {
	// test with a semantic token
	var tokens = TokenList{}
	var tokenType TokenType = SemX2lktSwitch
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: []*SemanticTokenType{},
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundToNoneOrLeft := token.IsBoundToOneOf(Bindings{BindNone, BindLeft})
	if !boundToNoneOrLeft {
		t.Errorf("token should be bound BindNone or BindLeft")
	}
	boundToUnknownOrRight := token.IsBoundToOneOf(Bindings{BindUnknown, BindRight})
	if boundToUnknownOrRight {
		t.Errorf("token should not be bound BindUnknown or BindRight")
	}
}

func TestToken_IsBoundToSem(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = SemX2lktSwitch
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: []*SemanticTokenType{},
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundToNone := token.IsBoundTo(BindNone)
	if !boundToNone {
		t.Errorf("token should be bound BindNone")
	}
	boundToRight := token.IsBoundTo(BindRight)
	if boundToRight {
		t.Errorf("token should not be bound BindRight")
	}
}

func TestToken_IsBoundToCF(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = CfEndOfOptions
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: CfEndOfOptions.SemanticCandidates,
		boundTo:            nil,
		value:              "--",
	}
	tokens = append(tokens, &token)
	boundToNone := token.IsBoundTo(BindNone)
	if !boundToNone {
		t.Errorf("token should be bound BindNone")
	}
	boundToLeft := token.IsBoundTo(BindLeft)
	if boundToLeft {
		t.Errorf("token should not be bound BindLeft")
	}
}

func TestToken_IsOption(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = CfOneDashWord
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: CfOneDashWord.SemanticCandidates,
		boundTo:            nil,
		value:              "-test",
	}
	tokens = append(tokens, &token)
	isOption := token.IsOptionPart()
	if !isOption {
		t.Errorf("token should be option")
	}
}
