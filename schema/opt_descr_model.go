package schema

type OptDescriptionModel []*OptDescription

func (models OptDescriptionModel) MatchArgument(arg string) []*SemanticTokenType {
	for _, description := range models {
		ttype := description.Match(arg)
		if ttype != nil {
			return ttype
		}
	}
	return nil
}
