package schema

type OptDescriptionModel []*OptDescription

// MatchArgument returns true if the provided argument matches at least one description
func (model OptDescriptionModel) MatchArgument(arg string) ([]*SemanticTokenType, bool) {
	for _, description := range model {
		ttype, matched := description.MatchArgument(arg)
		if matched {
			return ttype, true
		}
	}
	return nil, false
}

// Variants returns the option expression variants supported by each of its option description
func (model OptDescriptionModel) Variants() []*OptExpressionVariant {
	variantMap := map[*OptExpressionVariant]bool{}
	variants := make([]*OptExpressionVariant, 0, 10)
	for _, description := range model {
		for _, variant := range description.Variants() {
			if _, ok := variantMap[variant]; !ok {
				variantMap[variant] = true
				variants = append(variants, variant)
			}
		}
	}
	return variants
}
