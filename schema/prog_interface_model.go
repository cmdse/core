package schema

type OptDescriptionModel []*OptDescription

func (models OptDescriptionModel) MatchArgument(arg string) *SemanticTokenType {
	for _, description := range models {
		ttype := description.Match(arg)
		if ttype != nil {
			return ttype
		}
	}
	return nil
}

type ProgramInterfaceModel struct {
	optionScheme     OptionScheme
	descriptionModel OptDescriptionModel
}

func NewProgramInterfaceModel(optionScheme OptionScheme, optionDescriptions OptDescriptionModel) *ProgramInterfaceModel {
	return &ProgramInterfaceModel{
		optionScheme,
		optionDescriptions,
	}
}

func (pim *ProgramInterfaceModel) Scheme() OptionScheme {
	if pim == nil {
		return nil
	}
	return pim.optionScheme
}

func (pim *ProgramInterfaceModel) Descriptions() OptDescriptionModel {
	if pim == nil {
		return nil
	}
	return pim.descriptionModel
}
