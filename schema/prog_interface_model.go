package schema

type ProgramInterfaceModel struct {
	optionScheme       OptionScheme
	optionDescriptions []OptDescription
}

func NewProgramInterfaceModel(optionScheme OptionScheme, optionDescriptions []OptDescription) *ProgramInterfaceModel {
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

func (pim *ProgramInterfaceModel) Descriptions() []OptDescription {
	if pim == nil {
		return nil
	}
	return pim.optionDescriptions
}
