package parser

import (
	"github.com/cmdse/core/argparse/tkn"
)

// Behavior for running argument parsing
var ArgParseBehaviour = &Behavior{
	RunInferences: func(p *Parser, token *tkn.Token) {
		token.InferRight()
		token.InferLeft()
		token.InferPositional()
	},
	RunStaticChecks: func(p *Parser) {
		tokens := p.tokens
		tokens.CheckEndOfOptions()
		tokens.ReduceCandidatesWithScheme(p.pim.Scheme())
		tokens.MatchOptionDescription(p.pim.DescriptionModel())
	},
}
