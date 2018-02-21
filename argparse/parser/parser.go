package parser

import (
	"github.com/cmdse/core/schema"
)

// A Parser is an entity capable of turning an list of context-free tokens
// to a list of semantic tokens.
//
// See also
//
// * ParseTokens method
type Parser struct {
	*Behavior
	pim                 *schema.ProgramInterfaceModel
	tokens              schema.TokenList
	previousConversions int
	conversions         int
}

// Tokens return the list of tokens being processed for parsing
func (p *Parser) Tokens() schema.TokenList {
	return p.tokens
}

// PIM return the schema.ProgramInterfaceModel associated with the parsed program's arguments
func (p *Parser) PIM() *schema.ProgramInterfaceModel {
	return p.pim
}

func (p *Parser) lastTwoLoopsResultInConversion() bool {
	return p.previousConversions != 0 && p.conversions != 0
}

func (p *Parser) onePass() {
	p.previousConversions = p.conversions
	p.conversions = 0
	for _, token := range p.tokens.WhenContextFree() {
		p.RunInferences(p, token)
		if token.IsSemantic() {
			p.conversions++
		}
	}
}

// ParseTokens will turn an list of context-free tokens
// to a list of semantic tokens, when possible.
// The details of how it is done is encapsulated in the Behavior field.
//
// See also
//
// * Behavior
func (p *Parser) ParseTokens() schema.TokenList {
	p.RunStaticChecks(p)
	for {
		p.onePass()
		if !p.lastTwoLoopsResultInConversion() {
			return p.tokens
		}
	}
}
