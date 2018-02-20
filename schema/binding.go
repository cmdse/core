package schema

type Binding int

var bindingNames = map[Binding]string{
	BindUnknown: "UNKNOWN",
	BindNone:    "NONE",
	BindLeft:    "LEFT",
	BindRight:   "RIGHT",
}

const (
	BindUnknown Binding = iota
	BindNone
	BindLeft
	BindRight
)

func (binding Binding) String() string {
	val, ok := bindingNames[binding]
	if ok {
		return val
	} else {
		return ""
	}
}
