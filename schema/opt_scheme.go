package schema

type OptionScheme []*OptExpressionVariant

func (scheme OptionScheme) SupportsTokenType(tokenType *SemanticTokenType) bool {
	variant := tokenType.variant
	for _, testVariant := range scheme {
		if testVariant == variant {
			return true
		}
	}
	return false
}

var (
	OptionSchemePOSIXStrict = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXStackedShortSwitches,
		VariantPOSIXShortAssignment,
	}
	OptSchemeLinuxStandard = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXStackedShortSwitches,
		VariantPOSIXShortAssignment,
		VariantGNUSwitch,
		VariantGNUImplicitAssignment,
		VariantGNUExplicitAssignment,
		VariantEndOfOptions,
	}
	OptSchemeLinuxExplicit = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXStackedShortSwitches,
		VariantPOSIXShortAssignment,
		VariantGNUSwitch,
		VariantGNUExplicitAssignment,
		VariantEndOfOptions,
	}
	OptSchemeLinuxImplicit = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXStackedShortSwitches,
		VariantPOSIXShortAssignment,
		VariantGNUSwitch,
		VariantGNUImplicitAssignment,
		VariantEndOfOptions,
	}
	OptSchemeXToolkitStrict = OptionScheme{
		VariantX2lktImplicitAssignment,
		VariantX2lktExplicitAssignment,
		VariantX2lktReverseSwitch,
		VariantX2lktSwitch,
		VariantEndOfOptions,
	}
	OptSchemeXToolkitStandard = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXShortAssignment,
		VariantX2lktImplicitAssignment,
		VariantX2lktExplicitAssignment,
		VariantX2lktReverseSwitch,
		VariantX2lktSwitch,
		VariantEndOfOptions,
	}
	OptSchemeXToolkitExplicit = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantX2lktExplicitAssignment,
		VariantX2lktReverseSwitch,
		VariantX2lktSwitch,
		VariantEndOfOptions,
	}
	OptSchemeXToolkitImplicit = OptionScheme{
		VariantPOSIXShortSwitch,
		VariantPOSIXShortAssignment,
		VariantX2lktImplicitAssignment,
		VariantX2lktReverseSwitch,
		VariantX2lktSwitch,
		VariantEndOfOptions,
	}
)
