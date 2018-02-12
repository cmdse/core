package schema

import "regexp"

type MatchModel struct {
	variant *OptExpressionVariant
	flag    string
	// param is optional, zero value is the empty string
	param string
	regex *regexp.Regexp
}

func (matchModel *MatchModel) build() {
	matchModel.regex = matchModel.variant.Build(matchModel.flag, matchModel.param)
}

func NewSimpleMatchModel(variant *OptExpressionVariant, flag string) *MatchModel {
	matchModel := &MatchModel{
		variant,
		flag,
		"",
		nil,
	}
	matchModel.build()
	return matchModel
}

func NewMatchModelWithTypedValue(variant *OptExpressionVariant, flag string, param string) *MatchModel {
	matchModel := &MatchModel{
		variant,
		flag,
		param,
		nil,
	}
	matchModel.build()
	return matchModel
}

type OptDescription struct {
	Description string
	MatchModels []*MatchModel
}

// This function returns the semantic token type associated with the provided argument
// if it matched one of the MatchModel of this OptDescription,
// nil otherwise
func (optDescription *OptDescription) Match(arg string) *SemanticTokenType {
	for _, matchModel := range optDescription.MatchModels {
		if matchModel.regex.MatchString(arg) {
			return matchModel.variant.flagTokenType
		}
	}
	return nil
}
