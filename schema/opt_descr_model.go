package schema

type OptDescriptionModel []*OptDescription

func (model OptDescriptionModel) MatchArgument(arg string) ([]*SemanticTokenType, bool) {
	for _, description := range model {
		ttype, matched := description.MatchArgument(arg)
		if matched {
			return ttype, true
		}
	}
	return nil, false
}

func (model OptDescriptionModel) Variants() []*OptExpressionVariant {
	variantMap := map[*OptExpressionVariant]bool{}
	variants := make([]*OptExpressionVariant, 0, 10)
	for _, description := range model {
		for _, matchModel := range description.MatchModels {
			if _, ok := variantMap[matchModel.variant]; !ok {
				variantMap[matchModel.variant] = true
				variants = append(variants, matchModel.variant)
			}
		}
	}
	return variants
}
