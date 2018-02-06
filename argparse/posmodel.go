package argparse

type PositionalModel struct {
	Binding    Binding
	IsSemantic bool
	IsOption   bool
	name       string
}

func (posModel PositionalModel) String() string {
	return posModel.name
}

func (posModel PositionalModel) Equal(comparedPosModel *PositionalModel) bool {
	return posModel.name == comparedPosModel.name
}

var (
	OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE = PositionalModel{
		Binding:    RIGHT,
		IsSemantic: true,
		IsOption:   true,
		name:       "OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE",
	}
	OPT_IMPLICIT_ASSIGNMENT_VALUE = PositionalModel{
		Binding:    LEFT,
		IsSemantic: true,
		IsOption:   true,
		name:       "OPT_IMPLICIT_ASSIGNMENT_VALUE",
	}
	STANDALONE_OPT_ASSIGNMENT = PositionalModel{
		Binding:    NONE,
		IsSemantic: true,
		IsOption:   true,
		name:       "STANDALONE_OPT_ASSIGNMENT",
	}
	OPT_SWITCH = PositionalModel{
		Binding:    NONE,
		IsSemantic: true,
		IsOption:   true,
		name:       "OPT_SWITCH",
	}
	COMMAND_OPERAND = PositionalModel{
		Binding:    NONE,
		IsSemantic: true,
		IsOption:   false,
		name:       "COMMAND_OPERAND",
	}
	UNSET = PositionalModel{
		Binding:    UNKNOWN,
		IsSemantic: false,
		IsOption:   false,
		name:       "UNSET",
	}
)
