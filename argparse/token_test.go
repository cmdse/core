package argparse

import (
	. "cmdse-cli/schema"
	"testing"
)

func TestToken_IsBoundToOneOfCF(t *testing.T) {
	// Test with a context-free type
	var tokens = TokenList{}
	var tokenType TokenType = &CfOneDashWord
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: CfOneDashWord.SemanticCandidates,
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundNoneOrRight := token.IsBoundToOneOf(Bindings{NONE, RIGHT})
	if !boundNoneOrRight {
		t.Errorf("token should be bound NONE or RIGHT")
	}

	boundNoneOrLeft := token.IsBoundToOneOf(Bindings{NONE, LEFT})
	if boundNoneOrLeft {
		t.Errorf("token should not be bound NONE or LEFT")
	}
	boundUnknownOrLeft := token.IsBoundToOneOf(Bindings{UNKNOWN, LEFT})
	if boundUnknownOrLeft {
		t.Errorf("token should not be bound UNKNOWN or LEFT")
	}
}

func TestToken_IsBoundToOneOfSem(t *testing.T) {
	// test with a semantic token
	var tokens = TokenList{}
	var tokenType TokenType = &SemX2lktSwitch
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: []*SemanticTokenType{},
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundToNoneOrLeft := token.IsBoundToOneOf(Bindings{NONE, LEFT})
	if !boundToNoneOrLeft {
		t.Errorf("token should be bound NONE or LEFT")
	}
	boundToUnknownOrRight := token.IsBoundToOneOf(Bindings{UNKNOWN, RIGHT})
	if boundToUnknownOrRight {
		t.Errorf("token should not be bound UNKNOWN or RIGHT")
	}
}

func TestToken_IsBoundToSem(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = &SemX2lktSwitch
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: []*SemanticTokenType{},
		boundTo:            nil,
		value:              "-option",
	}
	tokens = append(tokens, &token)
	boundToNone := token.IsBoundTo(NONE)
	if !boundToNone {
		t.Errorf("token should be bound NONE")
	}
	boundToRight := token.IsBoundTo(RIGHT)
	if boundToRight {
		t.Errorf("token should not be bound RIGHT")
	}
}

func TestToken_IsBoundToCF(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = &CfEndOfOptions
	token := Token{
		argumentPosition:   0,
		ttype:              tokenType,
		tokens:             tokens,
		semanticCandidates: CfEndOfOptions.SemanticCandidates,
		boundTo:            nil,
		value:              "--",
	}
	tokens = append(tokens, &token)
	boundToNone := token.IsBoundTo(NONE)
	if !boundToNone {
		t.Errorf("token should be bound NONE")
	}
	boundToLeft := token.IsBoundTo(LEFT)
	if boundToLeft {
		t.Errorf("token should not be bound LEFT")
	}
}

func TestToken_IsOption(t *testing.T) {
	var tokens = TokenList{}
	var tokenType TokenType = &CfOneDashWord
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
