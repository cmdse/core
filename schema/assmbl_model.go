package schema

import "regexp"

type ExprAssemblyType int

const (
	AssmbTypeSplit     ExprAssemblyType = iota // tokens 1, options 1, use regex to split flag and value
	AssmbTypeGroup                             // tokens 2, options 1, assign left side to flag, right side to value
	AssmbTypeFlag                              // tokens 1, options 1, standalone flag
	AssmbTypeFlagStack                         // tokens 1, options n, split each letter to one AssmbTypeFlag
)

type OptionParts interface {
	Args() []string
}

type ExprAssemblyModel struct {
	atype    ExprAssemblyType
	Assemble func(OptionParts, *regexp.Regexp) *OptionExpression
}

var (
	AssmbModelSplit = &ExprAssemblyModel{
		AssmbTypeSplit,
		func(optionGroup OptionParts, regex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'Split' needs exactly one argument.")
			}
			groups := regex.FindStringSubmatch(args[0])[1:]
			if len(groups) != 2 {
				panic("Assembly model 'Split' regex has not exactly two matching groups")
			}
			return NewOptionExpression(OptionDefinition{
				&groups[0],
				&groups[1],
			})
		},
	}
	AssmbModelGroup = &ExprAssemblyModel{
		AssmbTypeGroup,
		func(optionGroup OptionParts, regex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 2 {
				panic("Assembly model 'Group' needs exactly two argument.")
			}
			return NewOptionExpression(OptionDefinition{
				&args[0],
				&args[1],
			})
		},
	}
	AssmbModelFlag = &ExprAssemblyModel{
		AssmbTypeFlag,
		func(optionGroup OptionParts, regex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'Flag' needs exactly one argument.")
			}
			return NewOptionExpression(OptionDefinition{
				&args[0],
				nil,
			})
		},
	}
	AssmbModelFlagStack = &ExprAssemblyModel{
		AssmbTypeFlagStack,
		func(optionGroup OptionParts, regex *regexp.Regexp) *OptionExpression {
			args := optionGroup.Args()
			if len(args) != 1 {
				panic("Assembly model 'Flag' needs exactly one argument.")
			}
			matchRes := regex.FindStringSubmatch(args[0])[1:]
			if len(matchRes) < 2 {
				panic("Assembly model 'FlagStack' regex has less then two matching groups")
			}
			options := make([]OptionDefinition, len(matchRes))
			for i := range matchRes {
				options[i] = OptionDefinition{&matchRes[i], nil}
			}
			return NewOptionExpression(options...)
		},
	}
)
