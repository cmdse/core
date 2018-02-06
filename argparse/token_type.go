package argparse

import "fmt"

type TokenType interface {
	PosModel() *PositionalModel
	IsSemantic() bool
	Name() string
	Equal(TokenType) bool
	fmt.Stringer
}
