package parser

import (
	"github.com/cmdse/core/argparse/tkn"
	. "github.com/cmdse/core/schema"
)

type Parser struct {
	*Behavior
	pim                 *ProgramInterfaceModel
	tokens              tkn.TokenList
	previousConversions int
	conversions         int
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

func (p *Parser) ParseTokens() tkn.TokenList {
	p.RunStaticChecks(p)
	for {
		p.onePass()
		if !p.lastTwoLoopsResultInConversion() {
			break
		}
	}
	return p.tokens
}
