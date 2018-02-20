package schema

// A ProgramInterfaceModel describes the command line interface capabilities of a program.
type ProgramInterfaceModel struct {
	optionScheme     OptionScheme
	descriptionModel OptDescriptionModel
}

// Create a ProgramInterfaceModel
func NewProgramInterfaceModel(optionScheme OptionScheme, optionDescriptions OptDescriptionModel) *ProgramInterfaceModel {
	return &ProgramInterfaceModel{
		optionScheme,
		optionDescriptions,
	}
}

// Return the OptionScheme in nil-safe mode.
// An OptionScheme is a set of option expression variants supported by a program command line interface.
func (pim *ProgramInterfaceModel) Scheme() OptionScheme {
	if pim == nil {
		return nil
	}
	return pim.optionScheme
}

// Return the OptDescriptionModel in nil-safe mode.
// An option description model is a set of option descriptions, which are composed
// of a description text field and a collection of match models.
// Each match model is related to an option expression variant and has a one-or-two groups regular expression.
// When two groups can be matched, the latest is the option parameter of an explicit option assignments.
func (pim *ProgramInterfaceModel) DescriptionModel() OptDescriptionModel {
	if pim == nil {
		return nil
	}
	return pim.descriptionModel
}
