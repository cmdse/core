package schema

type OptionStyle int

const (
	OptStyleNone OptionStyle = iota
	OptStyleXToolkit
	OptStyleGNU
	OptStylePOSIX
	OptStyleOld
)
