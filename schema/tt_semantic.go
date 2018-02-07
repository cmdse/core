package schema

type SemanticTokenType struct {
	posModel *PositionalModel
	name     string
	style    OptionStyle
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
		style:    OptStylePOSIX,
	}
	SemPosixStackedShortSwitches = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemPosixStackedShortSwitches",
		style:    OptStylePOSIX,
	}
	SemPosixShortAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemPosixShortAssignmentLeftSide",
		style:    OptStylePOSIX,
	}
	SemPosixShortAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemPosixShortAssignmentValue",
		style:    OptStylePOSIX,
	}
	SemPosixShortStickyValue = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemPosixShortStickyValue",
		style:    OptStylePOSIX,
	}
	// GNU
	SemGnuSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemGnuSwitch",
		style:    OptStyleGNU,
	}
	SemGnuExplicitAssignment = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemGnuExplicitAssignment",
		style:    OptStyleGNU,
	}
	SemGnuImplicitAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemGnuImplicitAssignmentLeftSide",
		style:    OptStyleGNU,
	}
	SemGnuImplicitAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemGnuImplicitAssignmentValue",
		style:    OptStyleGNU,
	}
	// X-Toolkit
	SemX2lktSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemX2lktSwitch",
		style:    OptStyleXToolkit,
	}
	SemX2lktReverseSwitch = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemX2lktReverseSwitch",
		style:    OptStyleXToolkit,
	}
	SemX2lktExplicitAssignment = SemanticTokenType{
		posModel: &STANDALONE_OPT_ASSIGNMENT,
		name:     "SemX2lktExplicitAssignment",
		style:    OptStyleXToolkit,
	}
	SemX2lktImplicitAssignmentLeftSide = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_LEFT_SIDE,
		name:     "SemX2lktImplicitAssignmentLeftSide",
		style:    OptStyleXToolkit,
	}
	SemX2lktImplicitAssignmentValue = SemanticTokenType{
		posModel: &OPT_IMPLICIT_ASSIGNMENT_VALUE,
		name:     "SemX2lktImplicitAssignmentValue",
		style:    OptStyleXToolkit,
	}
	// Special tokens
	SemEndOfOptions = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemEndOfOptions",
		style:    OptStyleNone,
	}
	SemOperand = SemanticTokenType{
		posModel: &COMMAND_OPERAND,
		name:     "SemOperand",
		style:    OptStyleNone,
	}
	SemHeadlessOption = SemanticTokenType{
		posModel: &OPT_SWITCH,
		name:     "SemHeadlessOption",
		style:    OptStyleOld,
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
