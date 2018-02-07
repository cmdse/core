package schema

type ExprAssemblyModel int

const (
	AssmbModelSplit     ExprAssemblyModel = iota // tokens 1, options 1, use regex to split flag and value
	AssmbModelGroup                              // tokens 2, options 1, assign left side to flag, right side to value
	AssmbModelFlag                               // tokens 1, options 1, standalone flag
	AssmbModelFlagStack                          // tokens 1, options n, split each letter to one AssmbModelFlag
)

type ExpressionVariant struct {
	style         OptionStyle
	assemblyModel ExprAssemblyModel
}

var (
	VarPOSIXShortSwitch = ExpressionVariant{
		style:         OptStylePOSIX,
		assemblyModel: AssmbModelFlag,
	}
	VarPOSIXStackedShortSwitches = ExpressionVariant{
		style:         OptStylePOSIX,
		assemblyModel: AssmbModelFlagStack,
	}
	VarPOSIXShortAssignment = ExpressionVariant{
		style:         OptStylePOSIX,
		assemblyModel: AssmbModelGroup,
	}
	VarPOSIXShortStickyValue = ExpressionVariant{
		style:         OptStylePOSIX,
		assemblyModel: AssmbModelSplit,
	}
	VarX2lktSwitch = ExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyModel: AssmbModelFlag,
	}
	VarX2lktReverseSwitch = ExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyModel: AssmbModelFlag,
	}
	VarX2lktImplicitAssignment = ExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyModel: AssmbModelGroup,
	}
	VarX2lktExplicitAssignment = ExpressionVariant{
		style:         OptStyleXToolkit,
		assemblyModel: AssmbModelSplit,
	}
	VarGNUSwitch = ExpressionVariant{
		style:         OptStyleGNU,
		assemblyModel: AssmbModelFlag,
	}
	VarGNUImplicitAssignment = ExpressionVariant{
		style:         OptStyleGNU,
		assemblyModel: AssmbModelGroup,
	}
	VarGNUExplicitAssignment = ExpressionVariant{
		style:         OptStyleGNU,
		assemblyModel: AssmbModelSplit,
	}
	VarEndOfOptions = ExpressionVariant{
		style:         OptStyleNone,
		assemblyModel: AssmbModelFlag,
	}
	VarHeadlessOption = ExpressionVariant{
		style:         OptStyleOld,
		assemblyModel: AssmbModelFlag,
	}
)
