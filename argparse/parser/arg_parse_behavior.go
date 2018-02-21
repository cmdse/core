package parser

import (
	"github.com/cmdse/core/schema"
)

// ArgParseBehavior is a behavior specifically designed for argument parsing.
var ArgParseBehavior = &Behavior{
	RunInferences: func(p *Parser, token *schema.Token) {
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
