package schema

import "fmt"

type TokenType interface {
	PosModel() *PositionalModel
	IsSemantic() bool
	Name() string
	Equal(TokenType) bool
	Variant() *OptExpressionVariant
	fmt.Stringer
}
