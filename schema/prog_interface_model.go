package schema

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

func (pim *ProgramInterfaceModel) DescriptionModel() OptDescriptionModel {
	if pim == nil {
		return nil
	}
	return pim.descriptionModel
}
