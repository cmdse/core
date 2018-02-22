package schema

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type OptExpressionVariant struct {
	style              OptionStyle
	assemblyRegex      *regexp.Regexp
	assemblyRegBuilder *ParametricRegexBuilder
	assemblyModel      *ExprAssemblyModel
	flagTokenType      *SemanticTokenType
	// can be nil when opt expression contains one token only
	optValueTokenType *SemanticTokenType
	name              string
}

// Assemble groups a token list of one or two tokens forming a "group part" to an option expression.
func (optVariant *OptExpressionVariant) Assemble(groupParts TokenList) (*OptionExpression, error) {
	if len(groupParts) < 1 {
		return nil, errors.New("assemble must receive a non-empty TokenList")
	}
	if groupParts[0].Ttype.Variant() != optVariant {
		return nil, fmt.Errorf("assemble must receive a TokenList which token type variants match %v variant", optVariant.name)
	}
	return optVariant.assemblyModel.Assemble(groupParts, optVariant)
}

func (optVariant *OptExpressionVariant) FlagTokenType() *SemanticTokenType {
	return optVariant.flagTokenType
}

func (optVariant *OptExpressionVariant) OptValueTokenType() *SemanticTokenType {
	return optVariant.optValueTokenType
}

func (optVariant *OptExpressionVariant) Name() string {
	return optVariant.name
}

func safeStringList(stringList []string) []string {
	buffer := make([]string, len(stringList))
	for i, str := range stringList {
		buffer[i] = regexp.QuoteMeta(str)
	}
	return buffer
}

// Build a leftSideRegex given a flagName name and a list paramAllowedValues.
// When paramAllowedValues is non-zero, it is evaluated as the concatenation of possible values a|b|c ... etc
func (optVariant *OptExpressionVariant) Build(flagName string, paramList []string) *regexp.Regexp {
	switch optVariant.assemblyModel.atype {
	case AssmbTypeFlagStack, AssmbTypeFlag:
		if len(paramList) > 0 {
			panic("Cannot give a paramList argument when Assembly Type is 'flag'")
		}
	}
	paramRegex := ""
	if len(paramList) > 0 {
		paramRegex = strings.Join(safeStringList(paramList), "|")
	}
	regex, err := optVariant.assemblyRegBuilder.Build(regexp.QuoteMeta(flagName), paramRegex)
	if err != nil {
		panic(err)
	}
	return regex
}

func NewOptExpressionVariant(style OptionStyle, builder *ParametricRegexBuilder, model *ExprAssemblyModel, name string) *OptExpressionVariant {
	return &OptExpressionVariant{
		style:              style,
		assemblyRegex:      builder.BuildDefault(),
		assemblyRegBuilder: builder,
		assemblyModel:      model,
		flagTokenType:      nil,
		optValueTokenType:  nil,
		name:               name,
	}
}

var (
	VariantPOSIXShortSwitch          = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashLetter, AssmbModelFlag, "VariantPOSIXShortSwitch")
	VariantPOSIXStackedShortSwitches = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashWordAlphaNum, AssmbModelFlagStack, "VariantPOSIXStackedShortSwitches")
	VariantPOSIXShortAssignment      = NewOptExpressionVariant(OptStylePOSIX, RegBuilderOneDashLetter, AssmbModelGroup, "VariantPOSIXShortAssignment")
	VariantPOSIXShortStickyValue     = NewOptExpressionVariant(OptStylePOSIX, RegBuilderPosixShortStickyValue, AssmbModelSplit, "VariantPOSIXShortStickyValue")
	VariantX2lktSwitch               = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderOneDashWord, AssmbModelFlag, "VariantX2lktSwitch")
	VariantX2lktReverseSwitch        = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderX2lktReverseSwitch, AssmbModelFlag, "VariantX2lktReverseSwitch")
	VariantX2lktImplicitAssignment   = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderOneDashWord, AssmbModelGroup, "VariantX2lktImplicitAssignment")
	VariantX2lktExplicitAssignment   = NewOptExpressionVariant(OptStyleXToolkit, RegBuilderX2lktExplicitAssignment, AssmbModelSplit, "VariantX2lktExplicitAssignment")
	VariantGNUSwitch                 = NewOptExpressionVariant(OptStyleGNU, RegBuilderTwoDashWord, AssmbModelFlag, "VariantGNUSwitch")
	VariantGNUImplicitAssignment     = NewOptExpressionVariant(OptStyleGNU, RegBuilderTwoDashWord, AssmbModelGroup, "VariantGNUImplicitAssignment")
	VariantGNUExplicitAssignment     = NewOptExpressionVariant(OptStyleGNU, RegBuilderGnuExplicitAssignment, AssmbModelSplit, "VariantGNUExplicitAssignment")
	VariantHeadlessOption            = NewOptExpressionVariant(OptStyleOld, RegBuilderOptWord, AssmbModelFlag, "VariantHeadlessOption")
	VariantEndOfOptions              = &OptExpressionVariant{
		style:         OptStyleNone,
		assemblyRegex: RegexEndOfOptions,
		assemblyModel: AssmModelNone,
		name:          "VariantEndOfOptions",
	}
)

var Variants = []*OptExpressionVariant{
	VariantPOSIXShortSwitch,
	VariantPOSIXStackedShortSwitches,
	VariantPOSIXShortAssignment,
	VariantPOSIXShortStickyValue,
	VariantX2lktSwitch,
	VariantX2lktReverseSwitch,
	VariantX2lktImplicitAssignment,
	VariantX2lktExplicitAssignment,
	VariantGNUSwitch,
	VariantGNUImplicitAssignment,
	VariantGNUExplicitAssignment,
	VariantHeadlessOption,
	VariantEndOfOptions,
}
