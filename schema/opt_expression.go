package schema

type OptionDefinition struct {
	flag            *string
	assignmentValue *string
}

type OptionExpression struct {
	options []OptionDefinition
}

func NewOptionExpression(options ...OptionDefinition) *OptionExpression {
	return &OptionExpression{
		options,
	}
}
