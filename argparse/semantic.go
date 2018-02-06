package argparse

type SemanticTokenType struct {
	posModel *PositionalModel
	name     string
}

func (tokenType SemanticTokenType) IsSemantic() bool {
	return tokenType.PosModel().IsSemantic
}

func (tokenType SemanticTokenType) PosModel() *PositionalModel {
	return tokenType.posModel
}

func (tokenType SemanticTokenType) Name() string {
	return tokenType.name
}

func (tokenType SemanticTokenType) String() string {
	return tokenType.name
}

func (tokenType SemanticTokenType) Equal(comparedTType TokenType) bool {
	return tokenType.Name() == comparedTType.Name()
}

var (
	X2LKT_REVERSE_SWITCH = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "X2LKT_REVERSE_SWITCH",
	}
	GNU_EXPLICIT_ASSIGNMENT = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "GNU_EXPLICIT_ASSIGNMENT",
	}
	X2LKT_EXPLICIT_ASSIGNMENT = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "X2LKT_EXPLICIT_ASSIGNMENT",
	}
	END_OF_OPTIONS = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "END_OF_OPTIONS",
	}
	POSIX_SHORT_SWITCH = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "POSIX_SHORT_SWITCH",
	}
	POSIX_STACKED_SHORT_SWITCHES = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "POSIX_STACKED_SHORT_SWITCHES",
	}
	POSIX_SHORT_ASSIGNMENT_LEFT_SIDE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "POSIX_SHORT_ASSIGNMENT_LEFT_SIDE",
	}
	POSIX_SHORT_ASSIGNMENT_VALUE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "POSIX_SHORT_ASSIGNMENT_VALUE",
	}
	POSIX_SHORT_STICKY_VALUE = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "POSIX_SHORT_STICKY_VALUE",
	}
	GNU_SWITCH = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "GNU_SWITCH",
	}
	GNU_IMPLICIT_ASSIGNMENT_LEFT_SIDE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "GNU_IMPLICIT_ASSIGNMENT_LEFT_SIDE",
	}
	GNU_IMPLICIT_ASSIGNMENT_VALUE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "GNU_IMPLICIT_ASSIGNMENT_VALUE",
	}
	X2LKT_SWITCH = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "X2LKT_SWITCH",
	}
	X2LKT_IMPLICIT_ASSIGNEMNT_LEFT_SIDE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "X2LKT_IMPLICIT_ASSIGNEMNT_LEFT_SIDE",
	}
	X2LKT_IMPLICIT_ASSIGNMENT_VALUE = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "X2LKT_IMPLICIT_ASSIGNMENT_VALUE",
	}
	OPERAND = SemanticTokenType{
		posModel: &COMMAND_OPERAND,
		name:     "OPERAND",
	}
	HEADLESS_OPTION = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "HEADLESS_OPTION",
	}
)

var SemanticTokenTypes = []*SemanticTokenType{
	&END_OF_OPTIONS,
	&GNU_EXPLICIT_ASSIGNMENT,
	&GNU_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
	&GNU_IMPLICIT_ASSIGNMENT_VALUE,
	&GNU_SWITCH,
	&X2LKT_SWITCH,
	&X2LKT_REVERSE_SWITCH,
	&X2LKT_EXPLICIT_ASSIGNMENT,
	&X2LKT_IMPLICIT_ASSIGNEMNT_LEFT_SIDE,
	&X2LKT_IMPLICIT_ASSIGNMENT_VALUE,
	&POSIX_SHORT_ASSIGNMENT_LEFT_SIDE,
	&POSIX_SHORT_ASSIGNMENT_VALUE,
	&POSIX_SHORT_STICKY_VALUE,
	&POSIX_SHORT_SWITCH,
	&POSIX_STACKED_SHORT_SWITCHES,
	&HEADLESS_OPTION,
	&OPERAND,
}
