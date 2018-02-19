package schema

type OptDescription struct {
	Description string
	MatchModels MatchModels
}

// This function returns the semantic tokens type associated with the provided argument
// if it matched at least one, nil otherwise
func (optDescription *OptDescription) Match(arg string) []*SemanticTokenType {
	var matches []*SemanticTokenType
	for _, matchModel := range optDescription.MatchModels {
		if matchModel.regex.MatchString(arg) {
			matches = append(matches, matchModel.variant.flagTokenType)
		}
	}
	return matches
}
