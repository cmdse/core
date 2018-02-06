package argparse

type TokenType interface {
	PosModel() *PositionalModel
	IsSemantic() bool
	Name() string
	String() string
}
