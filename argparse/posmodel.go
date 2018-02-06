package argparse

type Binding int

const (
	UNKNOWN Binding = iota
	NONE
	LEFT
	RIGHT
)

func (binding Binding) String() string {
	switch binding {
	case UNKNOWN:
		return "UNKNOWN"
	case NONE:
		return "NONE"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return ""
	}
}

type Bindings []Binding

func (bindings Bindings) Contains(bindingToCheck Binding) bool {
	for _, binding := range bindings {
		if binding == bindingToCheck {
			return true
		}
	}
	return false
}

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
