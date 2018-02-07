package argparse

type SemanticTokenType struct {
	posModel *PositionalModel
	name     string
}

func (tokenType *SemanticTokenType) IsSemantic() bool {
	return tokenType.PosModel().IsSemantic
}

func (tokenType *SemanticTokenType) PosModel() *PositionalModel {
	return tokenType.posModel
}

func (tokenType *SemanticTokenType) Name() string {
	return tokenType.name
}

func (tokenType *SemanticTokenType) String() string {
	return tokenType.name
}

func (tokenType *SemanticTokenType) Equal(comparedTType TokenType) bool {
	return tokenType.Name() == comparedTType.Name()
}

var (
	// Posix
	SemPosixShortSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemPosixShortSwitch",
	}
	SemPosixStackedShortSwitches = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemPosixStackedShortSwitches",
	}
	SemPosixShortAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemPosixShortAssignmentLeftSide",
	}
	SemPosixShortAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemPosixShortAssignmentValue",
	}
	SemPosixShortStickyValue = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemPosixShortStickyValue",
	}
	// GNU
	SemGnuSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemGnuSwitch",
	}
	SemGnuExplicitAssignment = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemGnuExplicitAssignment",
	}
	SemGnuImplicitAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemGnuImplicitAssignmentLeftSide",
	}
	SemGnuImplicitAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemGnuImplicitAssignmentValue",
	}
	// X-Toolkit
	SemX2lktSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemX2lktSwitch",
	}
	SemX2lktReverseSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemX2lktReverseSwitch",
	}
	SemX2lktExplicitAssignment = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemX2lktExplicitAssignment",
	}
	SemX2lktImplicitAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemX2lktImplicitAssignmentLeftSide",
	}
	SemX2lktImplicitAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemX2lktImplicitAssignmentValue",
	}
	// Special tokens
	SemEndOfOptions = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemEndOfOptions",
	}
	SemOperand = SemanticTokenType{
		posModel: &COMMAND_OPERAND,
		name:     "SemOperand",
	}
	SemHeadlessOption = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemHeadlessOption",
	}
)

var SemanticTokenTypes = []*SemanticTokenType{
	&SemEndOfOptions,
	&SemGnuExplicitAssignment,
	&SemGnuImplicitAssignmentLeftSide,
	&SemGnuImplicitAssignmentValue,
	&SemGnuSwitch,
	&SemX2lktSwitch,
	&SemX2lktReverseSwitch,
	&SemX2lktExplicitAssignment,
	&SemX2lktImplicitAssignmentLeftSide,
	&SemX2lktImplicitAssignmentValue,
	&SemPosixShortAssignmentLeftSide,
	&SemPosixShortAssignmentValue,
	&SemPosixShortStickyValue,
	&SemPosixShortSwitch,
	&SemPosixStackedShortSwitches,
	&SemHeadlessOption,
	&SemOperand,
}
