package schema

import (
	"errors"
	"fmt"
	"strings"
)

// The assembling mechanism to create option expressions from tokens
type ExprAssemblyType int

const (
	AssmbTypeSplit     ExprAssemblyType = iota // tokens 1, options 1, use leftSideRegex to split flagName and value
	AssmbTypeGroup                             // tokens 2, options 1, assign left side to flagName, right side to value
	AssmbTypeFlag                              // tokens 1, options 1, standalone flagName
	AssmbTypeFlagStack                         // tokens 1, options n, split each letter to one AssmbTypeFlag
	AssmbTypeNone                              // tokens 1, options 1
)

type OptionParts interface {
	Args() []string
}

// An implementation of the assembling mechanism to create option expressions from tokens
type ExprAssemblyModel struct {
	atype    ExprAssemblyType
	Assemble func(TokenList, *OptExpressionVariant) (*OptionExpression, error)
}

func matchLeftSideArg(variant *OptExpressionVariant, groupParts TokenList, expectedLen int, modelName string) ([]string, error) {
	leftPart := groupParts[0].Value
	matchRes := variant.assemblyRegex.FindStringSubmatch(leftPart)
	if len(matchRes) == 0 {
		return nil, fmt.Errorf("variant '%v' could not match left part '%v' with assembly model %v", variant.Name(), leftPart, modelName)
	}
	matchRes = matchRes[1:]
	if expectedLen != len(matchRes) {
		return nil, fmt.Errorf("assembly model '%v' leftSideRegex should have matched %v groups but matched %v instead", modelName, expectedLen, len(matchRes))
	}
	return matchRes, nil
}

var (
	// Split the token in two parts to build up an expression, the option flag and the assignment value.
	// The leftmost the option flag, and the rightmost the option assignment value.
	// nolint: dupl
	AssmbModelSplit = &ExprAssemblyModel{
		AssmbTypeSplit,
		func(groupParts TokenList, variant *OptExpressionVariant) (*OptionExpression, error) {
			if len(groupParts) != 1 {
				return nil, errors.New("assembly model 'Split' needs exactly one argument")
			}
			groups, err := matchLeftSideArg(variant, groupParts, 2, "Split")
			if err != nil {
				return nil, err
			}
			return NewOptionExpression(&OptionDefinition{
				variant,
				groups[0],
				groups[1],
			}), nil
		},
	}
	// Group two tokens in one expression.
	// The leftmost the option flag, and the rightmost the option assignment value.
	AssmbModelGroup = &ExprAssemblyModel{
		AssmbTypeGroup,
		func(groupParts TokenList, variant *OptExpressionVariant) (*OptionExpression, error) {
			if len(groupParts) != 2 {
				return nil, errors.New("assembly model 'Group' needs exactly two arguments")
			}
			groups, err := matchLeftSideArg(variant, groupParts, 1, "Group")
			if err != nil {
				return nil, err
			}
			return NewOptionExpression(&OptionDefinition{
				variant,
				groups[0],
				groupParts[1].Value,
			}), nil
		},
	}
	// Create an expression from one token, with no assignment value.
	AssmbModelFlag = &ExprAssemblyModel{
		AssmbTypeFlag,
		func(groupParts TokenList, variant *OptExpressionVariant) (*OptionExpression, error) {
			if len(groupParts) != 1 {
				return nil, errors.New("assembly model 'flag' needs exactly one argument")
			}
			groups, err := matchLeftSideArg(variant, groupParts, 1, "flag")
			if err != nil {
				return nil, err
			}
			return NewOptionExpression(&OptionDefinition{
				variant,
				groups[0],
				"",
			}), nil
		},
	}
	// Create an expression mapping to multiple options from one token, with no assignment value.
	AssmbModelFlagStack = &ExprAssemblyModel{
		AssmbTypeFlagStack,
		func(groupParts TokenList, variant *OptExpressionVariant) (*OptionExpression, error) {
			if len(groupParts) != 1 {
				return nil, errors.New("assembly model 'FlagStack' needs exactly one argument")
			}
			groups, err := matchLeftSideArg(variant, groupParts, 1, "FlagStack")
			if err != nil {
				return nil, err
			}
			flags := strings.Split(groups[0], "")
			options := make([]*OptionDefinition, len(flags))
			for i := range flags {
				options[i] = &OptionDefinition{variant, flags[i], ""}
			}
			return NewOptionExpression(options...), nil
		},
	}
	AssmModelNone = &ExprAssemblyModel{
		AssmbTypeNone,
		func(groupParts TokenList, variant *OptExpressionVariant) (*OptionExpression, error) {
			if len(groupParts) != 1 {
				return nil, errors.New("assembly model 'TypeNone' needs exactly one argument")
			}
			return NewOptionExpression(&OptionDefinition{variant, "", ""}), nil
		},
	}
)
