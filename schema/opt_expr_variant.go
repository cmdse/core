package schema

import "regexp"

type OptExpressionVariant struct {
	style         OptionStyle
	assemblyRegex *regexp.Regexp
	assemblyModel *ExprAssemblyModel
}

func (optVariant *OptExpressionVariant) Assemble(expression OptionParts) *OptionExpression {
	return optVariant.assemblyModel.Assemble(expression, optVariant.assemblyRegex)
}

var (
	VariantPOSIXShortSwitch = &OptExpressionVariant{
		style:         OptStylePOSIX,
		assemblyRegex: RegexOneDashLetter,
		assemblyModel: AssmbModelFlag,
	}
	VariantPOSIXStackedShortSwitches = &OptExpressionVariant{
		style:         OptStylePOSIX,
		assemblyRegex: RegexOneDashWordAlphaNum,
		assemblyModel: AssmbModelFlagStack,
	}
	VariantPOSIXShortAssignment = &OptExpressionVariant{
		style:         OptStylePOSIX,
		assemblyModel: AssmbModelGroup,
	}
	VariantPOSIXShortStickyValue = &OptExpressionVariant{
		style:         OptStylePOSIX,
		assemblyRegex: RegexPosixShortStickyValue,
		assemblyModel: AssmbModelSplit,
	}
	VariantX2lktSwitch = &OptExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyRegex: RegexOneDashWord,
		assemblyModel: AssmbModelFlag,
	}
	VariantX2lktReverseSwitch = &OptExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyRegex: RegexX2lktReverseSwitch,
		assemblyModel: AssmbModelFlag,
	}
	VariantX2lktImplicitAssignment = &OptExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyRegex: RegexX2lktReverseSwitch,
		assemblyModel: AssmbModelGroup,
	}
	VariantX2lktExplicitAssignment = &OptExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyRegex: RegexX2lktExplicitAssignment,
		assemblyModel: AssmbModelSplit,
	}
	VariantGNUSwitch = &OptExpressionVariant{
		style:         OptStyleGNU,
		assemblyRegex: RegexTwoDashWord,
		assemblyModel: AssmbModelFlag,
	}
	VariantGNUImplicitAssignment = &OptExpressionVariant{
		style:         OptStyleGNU,
		assemblyRegex: RegexTwoDashWord,
		assemblyModel: AssmbModelGroup,
	}
	VariantGNUExplicitAssignment = &OptExpressionVariant{
		style:         OptStyleGNU,
		assemblyRegex: RegexGnuExplicitAssignment,
		assemblyModel: AssmbModelSplit,
	}
	VariantEndOfOptions = &OptExpressionVariant{
		style:         OptStyleNone,
		assemblyRegex: RegexEndOfOptions,
		assemblyModel: AssmbModelFlag,
	}
	VariantHeadlessOption = &OptExpressionVariant{
		style:         OptStyleOld,
		assemblyRegex: RegexMatchAny,
		assemblyModel: AssmbModelFlag,
	}
)
