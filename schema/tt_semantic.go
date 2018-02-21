package schema

type SemanticTokenType struct {
	posModel *PositionalModel
	name     string
	// If is option part, the associated Variant
	// nil otherwise
	variant *OptExpressionVariant
	style   OptionStyle
}

func NewSemanticTokenType(posModel *PositionalModel, name string, variant *OptExpressionVariant, style OptionStyle) *SemanticTokenType {
	ttype := &SemanticTokenType{
		posModel,
		name,
		variant,
		style,
	}
	if variant != nil {
		if posModel.IsOptionFlag {
			variant.flagTokenType = ttype
		} else { // is Option Value
			variant.optValueTokenType = ttype
		}
	}
	return ttype
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
	SemPOSIXShortSwitch = NewSemanticTokenType(
		PosModOptSwitch,
		"SemPOSIXShortSwitch",
		VariantPOSIXShortSwitch,
		OptStylePOSIX,
	)
	SemPOSIXStackedShortSwitches = NewSemanticTokenType(
		PosModOptSwitch,
		"SemPOSIXStackedShortSwitches",
		VariantPOSIXStackedShortSwitches,
		OptStylePOSIX,
	)
	SemPOSIXShortAssignmentLeftSide = NewSemanticTokenType(
		PosModOptImplicitAssignmentLeftSide,
		"SemPOSIXShortAssignmentLeftSide",
		VariantPOSIXShortAssignment,
		OptStylePOSIX,
	)
	SemPOSIXShortAssignmentValue = NewSemanticTokenType(
		PosModOptImplicitAssignmentValue,
		"SemPOSIXShortAssignmentValue",
		VariantPOSIXShortAssignment,
		OptStylePOSIX,
	)
	SemPOSIXShortStickyValue = NewSemanticTokenType(
		PosModStandaloneOptAssignment,
		"SemPOSIXShortStickyValue",
		VariantPOSIXShortStickyValue,
		OptStylePOSIX,
	)
	// GNU
	SemGNUSwitch = NewSemanticTokenType(
		PosModOptSwitch,
		"SemGNUSwitch",
		VariantGNUSwitch,
		OptStyleGNU,
	)
	SemGNUExplicitAssignment = NewSemanticTokenType(
		PosModStandaloneOptAssignment,
		"SemGNUExplicitAssignment",
		VariantGNUExplicitAssignment,
		OptStyleGNU,
	)
	SemGNUImplicitAssignmentLeftSide = NewSemanticTokenType(
		PosModOptImplicitAssignmentLeftSide,
		"SemGNUImplicitAssignmentLeftSide",
		VariantGNUImplicitAssignment,
		OptStyleGNU,
	)
	SemGNUImplicitAssignmentValue = NewSemanticTokenType(
		PosModOptImplicitAssignmentValue,
		"SemGNUImplicitAssignmentValue",
		VariantGNUImplicitAssignment,
		OptStyleGNU,
	)
	// X-Toolkit
	SemX2lktSwitch = NewSemanticTokenType(
		PosModOptSwitch,
		"SemX2lktSwitch",
		VariantX2lktSwitch,
		OptStyleXToolkit,
	)
	SemX2lktReverseSwitch = NewSemanticTokenType(
		PosModOptSwitch,
		"SemX2lktReverseSwitch",
		VariantX2lktReverseSwitch,
		OptStyleXToolkit,
	)
	SemX2lktExplicitAssignment = NewSemanticTokenType(
		PosModStandaloneOptAssignment,
		"SemX2lktExplicitAssignment",
		VariantX2lktExplicitAssignment,
		OptStyleXToolkit,
	)
	SemX2lktImplicitAssignmentLeftSide = NewSemanticTokenType(
		PosModOptImplicitAssignmentLeftSide,
		"SemX2lktImplicitAssignmentLeftSide",
		VariantX2lktImplicitAssignment,
		OptStyleXToolkit,
	)
	SemX2lktImplicitAssignmentValue = NewSemanticTokenType(
		PosModOptImplicitAssignmentValue,
		"SemX2lktImplicitAssignmentValue",
		VariantX2lktImplicitAssignment,
		OptStyleXToolkit,
	)
	// Special tokens
	SemEndOfOptions = NewSemanticTokenType(
		PosModOptSwitch,
		"SemEndOfOptions",
		VariantEndOfOptions,
		OptStyleNone,
	)
	SemOperand = NewSemanticTokenType(
		PosModCommandOperand,
		"SemOperand",
		nil,
		OptStyleNone,
	)
	SemHeadlessOption = NewSemanticTokenType(
		PosModOptSwitch,
		"SemHeadlessOption",
		VariantHeadlessOption,
		OptStyleOld,
	)
)

var SemanticTokenTypes = []*SemanticTokenType{
	SemEndOfOptions,
	SemGNUExplicitAssignment,
	SemGNUImplicitAssignmentLeftSide,
	SemGNUImplicitAssignmentValue,
	SemGNUSwitch,
	SemX2lktSwitch,
	SemX2lktReverseSwitch,
	SemX2lktExplicitAssignment,
	SemX2lktImplicitAssignmentLeftSide,
	SemX2lktImplicitAssignmentValue,
	SemPOSIXShortAssignmentLeftSide,
	SemPOSIXShortAssignmentValue,
	SemPOSIXShortStickyValue,
	SemPOSIXShortSwitch,
	SemPOSIXStackedShortSwitches,
	SemHeadlessOption,
	SemOperand,
}
