package schema

type SemanticTokenType struct {
	posModel *PositionalModel
	name     string
	// If is option part, the associated variant
	// nil otherwise
	variant *OptExpressionVariant
	style   OptionStyle
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

func (tokenType *SemanticTokenType) Variant() *OptExpressionVariant {
	return tokenType.variant
}

func (tokenType *SemanticTokenType) Equal(comparedTType TokenType) bool {
	return tokenType.Name() == comparedTType.Name()
}

var (
	// Posix
	SemPosixShortSwitch = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemPosixShortSwitch",
		variant:  VariantPOSIXShortSwitch,
		style:    OptStylePOSIX,
	}
	SemPosixStackedShortSwitches = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemPosixStackedShortSwitches",
		variant:  VariantPOSIXStackedShortSwitches,
		style:    OptStylePOSIX,
	}
	SemPosixShortAssignmentLeftSide = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentLeftSide,
		name:     "SemPosixShortAssignmentLeftSide",
		variant:  VariantPOSIXShortAssignment,
		style:    OptStylePOSIX,
	}
	SemPosixShortAssignmentValue = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentValue,
		name:     "SemPosixShortAssignmentValue",
		variant:  VariantPOSIXShortAssignment,
		style:    OptStylePOSIX,
	}
	SemPosixShortStickyValue = &SemanticTokenType{
		posModel: PosModStandaloneOptAssignment,
		name:     "SemPosixShortStickyValue",
		variant:  VariantPOSIXShortStickyValue,
		style:    OptStylePOSIX,
	}
	// GNU
	SemGnuSwitch = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemGnuSwitch",
		variant:  VariantGNUSwitch,
		style:    OptStyleGNU,
	}
	SemGnuExplicitAssignment = &SemanticTokenType{
		posModel: PosModStandaloneOptAssignment,
		name:     "SemGnuExplicitAssignment",
		variant:  VariantGNUExplicitAssignment,
		style:    OptStyleGNU,
	}
	SemGnuImplicitAssignmentLeftSide = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentLeftSide,
		name:     "SemGnuImplicitAssignmentLeftSide",
		variant:  VariantGNUImplicitAssignment,
		style:    OptStyleGNU,
	}
	SemGnuImplicitAssignmentValue = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentValue,
		name:     "SemGnuImplicitAssignmentValue",
		variant:  VariantGNUImplicitAssignment,
		style:    OptStyleGNU,
	}
	// X-Toolkit
	SemX2lktSwitch = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemX2lktSwitch",
		variant:  VariantX2lktSwitch,
		style:    OptStyleXToolkit,
	}
	SemX2lktReverseSwitch = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemX2lktReverseSwitch",
		variant:  VariantX2lktReverseSwitch,
		style:    OptStyleXToolkit,
	}
	SemX2lktExplicitAssignment = &SemanticTokenType{
		posModel: PosModStandaloneOptAssignment,
		name:     "SemX2lktExplicitAssignment",
		variant:  VariantX2lktExplicitAssignment,
		style:    OptStyleXToolkit,
	}
	SemX2lktImplicitAssignmentLeftSide = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentLeftSide,
		name:     "SemX2lktImplicitAssignmentLeftSide",
		variant:  VariantX2lktImplicitAssignment,
		style:    OptStyleXToolkit,
	}
	SemX2lktImplicitAssignmentValue = &SemanticTokenType{
		posModel: PosModOptImplicitAssignmentValue,
		name:     "SemX2lktImplicitAssignmentValue",
		variant:  VariantX2lktImplicitAssignment,
		style:    OptStyleXToolkit,
	}
	// Special tokens
	SemEndOfOptions = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemEndOfOptions",
		variant:  VariantEndOfOptions,
		style:    OptStyleNone,
	}
	SemOperand = &SemanticTokenType{
		posModel: PosModCommandOperand,
		name:     "SemOperand",
		variant:  nil,
		style:    OptStyleNone,
	}
	SemHeadlessOption = &SemanticTokenType{
		posModel: PosModOptSwitch,
		name:     "SemHeadlessOption",
		variant:  VariantHeadlessOption,
		style:    OptStyleOld,
	}
)

var SemanticTokenTypes = []*SemanticTokenType{
	SemEndOfOptions,
	SemGnuExplicitAssignment,
	SemGnuImplicitAssignmentLeftSide,
	SemGnuImplicitAssignmentValue,
	SemGnuSwitch,
	SemX2lktSwitch,
	SemX2lktReverseSwitch,
	SemX2lktExplicitAssignment,
	SemX2lktImplicitAssignmentLeftSide,
	SemX2lktImplicitAssignmentValue,
	SemPosixShortAssignmentLeftSide,
	SemPosixShortAssignmentValue,
	SemPosixShortStickyValue,
	SemPosixShortSwitch,
	SemPosixStackedShortSwitches,
	SemHeadlessOption,
	SemOperand,
}
