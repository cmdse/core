package argparse

import (
	. "github.com/cmdse/core/schema"
)

type Parser struct {
	tokens              TokenList
	previousConversions int
	conversions         int
}

func newParser(tokens TokenList) *Parser {
	return &Parser{
		tokens,
		1,
		1,
	}
}

func parseArgument(arg string) *ContextFreeTokenType {
	var fetchedType *ContextFreeTokenType
	for _, ttype := range ContextFreeTokenTypes {
		regex, isMatchAll := ttype.Regexp()
		if isMatchAll {
			fetchedType = ttype
			break
		} else {
			if regex.MatchString(arg) {
				fetchedType = ttype
				break
			}
		}
	}
	return fetchedType
}

func initTokens(args []string) TokenList {
	tokens := make([]*Token, len(args))
	for i, arg := range args {
		contextFreeTType := parseArgument(arg)
		var tokenType TokenType = contextFreeTType
		token := newToken(i, tokenType, arg, tokens)
		tokens[i] = token
		token.possiblyConvertToSemantic()
	}
	return tokens
}

func (p *Parser) lastTwoLoopsResultInConversion() bool {
	return p.previousConversions != 0 && p.conversions != 0
}

func (p *Parser) onePass() {
	p.previousConversions = p.conversions
	p.conversions = 0
	for _, token := range p.tokens.whenContextFree() {
		token.InferRight()
		token.InferLeft()
		token.InferPositional()
		if token.IsSemantic() {
			p.conversions++
		}
	}
}

func (p *Parser) parseTokens(pim *ProgramInterfaceModel) TokenList {
	tokens := p.tokens
	tokens.CheckEndOfOptions()
	tokens.ReduceCandidatesWithScheme(pim.Scheme())
	tokens.MatchOptionDescription(pim.DescriptionModel())
	for {
		p.onePass()
		if !p.lastTwoLoopsResultInConversion() {
			break
		}
	}
	return tokens
}

func ParseArguments(args []string, programInterfaceModel *ProgramInterfaceModel) TokenList {
	tokens := initTokens(args)
	parser := newParser(tokens)
	return parser.parseTokens(programInterfaceModel)
}
