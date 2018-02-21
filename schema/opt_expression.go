package schema

type OptionDefinition struct {
	variant         *OptExpressionVariant
	flag            string
	assignmentValue string
}

func (optDefinition *OptionDefinition) Variant() *OptExpressionVariant {
	return optDefinition.variant
}

func (optDefinition *OptionDefinition) Flag() string {
	return optDefinition.flag
}

func (optDefinition *OptionDefinition) AssignmentValue() string {
	return optDefinition.assignmentValue
}

type OptionExpression struct {
	options []*OptionDefinition
}

func (optExpression *OptionExpression) Options() []*OptionDefinition {
	return optExpression.options
}

func NewOptionExpression(options ...*OptionDefinition) *OptionExpression {
	return &OptionExpression{
		options,
	}
}
