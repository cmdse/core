package schema

import "regexp"

type OptExpressionVariant struct {
	style              OptionStyle
	assemblyRegex      *regexp.Regexp
	assemblyRegBuilder *ParametricRegexBuilder
	assemblyModel      *ExprAssemblyModel
	flagTokenType      *SemanticTokenType
	// can be nil when opt expression contains one token only
	optValueTokenType *SemanticTokenType
}

func (optVariant *OptExpressionVariant) Assemble(expression OptionParts) *OptionExpression {
	return optVariant.assemblyModel.Assemble(expression, optVariant.assemblyRegex)
}

func (optVariant *OptExpressionVariant) FlagTokenType() *SemanticTokenType {
	return optVariant.flagTokenType
}

func (optVariant *OptExpressionVariant) OptValueTokenType() *SemanticTokenType {
	return optVariant.optValueTokenType
}

// Build a regex given a flag name and optional param
// empty string stands for no param, or default matcher
func (optVariant *OptExpressionVariant) Build(flag string, param string) *regexp.Regexp {
	switch optVariant.assemblyModel.atype {
	case AssmbTypeFlagStack, AssmbTypeFlag:
		if param != "" {
			panic("Cannot give a param argument when Assembly Type is 'Flag'")
		}
	}
	regex, err := optVariant.assemblyRegBuilder.Build(regexp.QuoteMeta(flag), param)
	if err != nil {
		panic(err)
	}
	return regex
}

func NewOptExpressionVariant(style OptionStyle, builder *ParametricRegexBuilder, model *ExprAssemblyModel) *OptExpressionVariant {
	return &OptExpressionVariant{
		style:              style,
		assemblyRegex:      builder.BuildDefault(),
		assemblyRegBuilder: builder,
		assemblyModel:      model,
		flagTokenType:      nil,
		optValueTokenType:  nil,
	}
}

var (
	VariantPOSIXShortSwitch          = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashLetter, AssmbModelFlag)
	VariantPOSIXStackedShortSwitches = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashWordAlphaNum, AssmbModelFlagStack)
	VariantPOSIXShortAssignment      = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashLetter, AssmbModelGroup)
	VariantPOSIXShortStickyValue     = NewOptExpressionVariant(OptStylePOSIX, RegBuilderPosixShortStickyValue, AssmbModelSplit)
	VariantX2lktSwitch               = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderOneDashWord, AssmbModelFlag)
	VariantX2lktReverseSwitch        = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderX2lktReverseSwitch, AssmbModelFlag)
	VariantX2lktImplicitAssignment   = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderX2lktReverseSwitch, AssmbModelGroup)
	VariantX2lktExplicitAssignment   = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderX2lktExplicitAssignment, AssmbModelSplit)
	VariantGNUSwitch                 = NewOptExpressionVariant(OptStyleGNU, RegBuilderTwoDashWord, AssmbModelFlag)
	VariantGNUImplicitAssignment     = NewOptExpressionVariant(OptStyleGNU, RegBuilderTwoDashWord, AssmbModelGroup)
	VariantGNUExplicitAssignment     = NewOptExpressionVariant(OptStyleGNU, RegBuilderGnuExplicitAssignment, AssmbModelSplit)
	VariantHeadlessOption            = NewOptExpressionVariant(OptStyleOld, RegBuilderOptWord, AssmbModelFlag)
	VariantEndOfOptions              = &OptExpressionVariant{
		style:         OptStyleNone,
		assemblyRegex: RegexEndOfOptions,
		assemblyModel: AssmbModelFlag,
	}
)
