package schema

type OptDescriptionModel []*OptDescription

func (model OptDescriptionModel) MatchArgument(arg string) []*SemanticTokenType {
	for _, description := range model {
		ttype := description.MatchArgument(arg)
		if ttype != nil {
			return ttype
		}
	}
	return nil
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
