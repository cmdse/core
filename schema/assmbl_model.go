package schema

import "regexp"

// The assembling mechanism to create option expressions from tokens
type ExprAssemblyType int

const (
	AssmbTypeSplit     ExprAssemblyType = iota // tokens 1, options 1, use leftSideRegex to split flagName and value
	AssmbTypeGroup                             // tokens 2, options 1, assign left side to flagName, right side to value
	AssmbTypeFlag                              // tokens 1, options 1, standalone flagName
	AssmbTypeFlagStack                         // tokens 1, options n, split each letter to one AssmbTypeFlag
)

type OptionParts interface {
	Args() []string
}

// An implementation of the assembling mechanism to create option expressions from tokens
type ExprAssemblyModel struct {
	atype    ExprAssemblyType
	Assemble func(OptionParts, *regexp.Regexp) *OptionExpression
}

var (
	// Split the token in two parts to build up an expression, the option flag and the assignment value.
	// The leftmost the option flag, and the rightmost the option assignment value.
	AssmbModelSplit = &ExprAssemblyModel{
		AssmbTypeSplit,
		func(optionGroup OptionParts, leftSideRegex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'Split' needs exactly one argument.")
			}
			groups := leftSideRegex.FindStringSubmatch(args[0])[1:]
			if len(groups) != 2 {
				panic("Assembly model 'Split' leftSideRegex has not exactly two matching groups")
			}
			return NewOptionExpression(OptionDefinition{
				groups[0],
				groups[1],
			})
		},
	}
	// Group two tokens in one expression.
	// The leftmost the option flag, and the rightmost the option assignment value.
	AssmbModelGroup = &ExprAssemblyModel{
		AssmbTypeGroup,
		func(optionGroup OptionParts, leftSideRegex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 2 {
				panic("Assembly model 'Group' needs exactly two argument.")
			}
			return NewOptionExpression(OptionDefinition{
				args[0],
				args[1],
			})
		},
	}
	// Create an expression from one token, with no assignment value.
	AssmbModelFlag = &ExprAssemblyModel{
		AssmbTypeFlag,
		func(optionGroup OptionParts, leftSideRegex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'Flag' needs exactly one argument.")
			}
			return NewOptionExpression(OptionDefinition{
				args[0],
				"",
			})
		},
	}
	// Create an expression mapping to multiple options from one token, with no assignment value.
	AssmbModelFlagStack = &ExprAssemblyModel{
		AssmbTypeFlagStack,
		func(optionGroup OptionParts, leftSideRegex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'FlagStack' needs exactly one argument.")
			}
			matchRes := leftSideRegex.FindStringSubmatch(args[0])[1:]
			if len(matchRes) < 2 {
				panic("Assembly model 'FlagStack' leftSideRegex has less then two matching groups")
			}
			options := make([]OptionDefinition, len(matchRes))
			for i := range matchRes {
				options[i] = OptionDefinition{matchRes[i], ""}
			}
			return NewOptionExpression(options...)
		},
	}
)
