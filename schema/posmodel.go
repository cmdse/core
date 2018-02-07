package schema

type PositionalModel struct {
	Binding      Binding
	IsSemantic   bool
	IsOptionPart bool
	IsOptionFlag bool
	name         string
}

func (posModel PositionalModel) String() string {
	return posModel.name
}

func (posModel PositionalModel) Equal(comparedPosModel *PositionalModel) bool {
	return posModel.name == comparedPosModel.name
}

var (
	PosModOptImplicitAssignmentLeftSide = &PositionalModel{
		Binding:      RIGHT,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModOptImplicitAssignmentLeftSide",
	}
	PosModOptImplicitAssignmentValue = &PositionalModel{
		Binding:      LEFT,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: false,
		name:         "PosModOptImplicitAssignmentValue",
	}
	PosModStandaloneOptAssignment = &PositionalModel{
		Binding:      NONE,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModStandaloneOptAssignment",
	}
	PosModOptSwitch = &PositionalModel{
		Binding:      NONE,
		IsSemantic:   true,
		IsOptionPart: true,
		IsOptionFlag: true,
		name:         "PosModOptSwitch",
	}
	PosModCommandOperand = &PositionalModel{
		Binding:      NONE,
		IsSemantic:   true,
		IsOptionPart: false,
		IsOptionFlag: false,
		name:         "PosModCommandOperand",
	}
	PosModUnset = &PositionalModel{
		Binding:      UNKNOWN,
		IsSemantic:   false,
		IsOptionPart: false,
		IsOptionFlag: false,
		name:         "PosModUnset",
	}
)
