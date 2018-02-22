package schema

import (
	"fmt"
	"regexp"
	"strings"
)

// MatchModel is a structure to recognize option expressions
//
// Internally, it is composed of a flagName, an optional paramName part and an OpExpressionVariant
// flagName is the option identifier, with any leading hyphens (or + sign) stripped
// paramName is optional, relevant for option assignments
//
// Example: --flagName=paramName, VariantGNUExplicitAssignment
//
//
type MatchModel struct {
	description        string
	variant            *OptExpressionVariant
	flagName           string
	paramName          string
	paramAllowedValues []string
	leftSideRegex      *regexp.Regexp
}

// A slice of MatchModel
type MatchModels []*MatchModel

func (matchModel *MatchModel) build() {
	matchModel.leftSideRegex = matchModel.variant.Build(matchModel.flagName, matchModel.paramAllowedValues)
}

func (matchModel *MatchModel) Variant() *OptExpressionVariant {
	return matchModel.variant
}

func (matchModel *MatchModel) FlagName() string {
	return matchModel.flagName
}

func (matchModel *MatchModel) ParamName() string {
	return matchModel.paramName
}

func (matchModel *MatchModel) ParamAllowedValues() []string {
	return matchModel.paramAllowedValues
}

func (matchModel *MatchModel) LeftSideRegex() *regexp.Regexp {
	return matchModel.leftSideRegex
}

// MatchLeftSide matches the left side of an option expression against it's model.
func (matchModel *MatchModel) MatchLeftSide(arg string) bool {
	return matchModel.leftSideRegex.MatchString(arg)
}

// Return the description bound to the underlying OptionDescription.
func (matchModel *MatchModel) Description() string {
	return matchModel.description
}

// NewStandaloneMatchModel creates a MatchModel given an OptExpressionVariant and a flagName.
// flagName is the option identifier, with any leading hyphens (or + sign) stripped
//
// Examples :
//
// MatchModel grabbing '--opt' :
//
//   NewStandaloneMatchModel(VariantGNUSwitch, "opt")
//
//
// MatchModel grabbing '-switch' :
//
//   NewStandaloneMatchModel(VariantX2lktSwitch, "switch")
//
//
// MatchModel grabbing '+switch' :
//
//   NewStandaloneMatchModel(VariantX2lktReverseSwitch, "switch")
//
//
// MatchModel grabbing '-p' :
//
//   NewStandaloneMatchModel(VariantPOSIXShortSwitch, "p")
//
func NewStandaloneMatchModel(variant *OptExpressionVariant, flagName string) *MatchModel {
	matchModel := &MatchModel{
		"",
		variant,
		flagName,
		"",
		nil,
		nil,
	}
	matchModel.build()
	return matchModel
}

// NewAssignmentMatchModel creates a MatchModel given an OptExpressionVariant, a flagName, a paramName.
// flagName is the option identifier, with any leading hyphens (or + sign) stripped
// paramName won't be used for matching and is just provided as a placeholder for information purpose.
//
// Examples :
//
// MatchModel grabbing '--opt=<something>' :
//
//   NewStandaloneMatchModel(VariantGNUExplicitAssignment, "opt", "something")
//
//
// MatchModel grabbing '--option <something>' :
//
//   NewStandaloneMatchModel(VariantGNUImplicitAssignment, "option", "something")
//
//
// MatchModel grabbing '-option=<something>' :
//
//   NewStandaloneMatchModel(VariantX2lktExplicitAssignment, "option", "something")
//
//
// MatchModel grabbing '-p<n>' :
//
//   NewStandaloneMatchModel(VariantPOSIXShortStickyValue, "p", "n")
//
func NewAssignmentMatchModel(variant *OptExpressionVariant, flagName string, paramName string) *MatchModel {
	matchModel := &MatchModel{
		"",
		variant,
		flagName,
		paramName,
		nil,
		nil,
	}
	matchModel.build()
	return matchModel
}

// NewMatchModelFromDefinition creates a MatchModel from an OptionDefinition
func NewMatchModelFromDefinition(description *OptionDefinition) *MatchModel {
	matchModel := &MatchModel{
		"",
		description.Variant(),
		description.Flag(),
		description.AssignmentValue(),
		nil,
		nil,
	}
	matchModel.build()
	return matchModel

}

func extractPOSIXShortSwitchFlags(model OptDescriptionModel) []string {
	posixShortSwitchModels := make([]string, 0, len(model))
	for _, descr := range model {
		for _, matchModel := range descr.MatchModels {
			if matchModel.variant == VariantPOSIXShortSwitch {
				posixShortSwitchModels = append(posixShortSwitchModels, matchModel.flagName)
			}
		}
	}
	return posixShortSwitchModels
}

// NewPOSIXStackMatchModel creates a MatchModel from an OptDescriptionModel
// Returned instance matches any expression composed of 'model' POSIXShortSwitch option variants.
func NewPOSIXStackMatchModel(model OptDescriptionModel) *MatchModel {
	posixShortSwitchFlags := extractPOSIXShortSwitchFlags(model)
	leftSideRegex := fmt.Sprintf("^-[%v]{2,}$", strings.Join(posixShortSwitchFlags, ","))
	matchModel := &MatchModel{
		"",
		VariantPOSIXShortSwitch,
		"",
		"",
		nil,
		regexp.MustCompile(leftSideRegex),
	}
	return matchModel
}
