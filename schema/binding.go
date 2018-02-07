package schema

type Binding int

const (
	BindUnknown Binding = iota
	BindNone
	BindLeft
	BindRight
)

func (binding Binding) String() string {
	switch binding {
	case BindUnknown:
		return "BindUnknown"
	case BindNone:
		return "BindNone"
	case BindLeft:
		return "BindLeft"
	case BindRight:
		return "BindRight"
	default:
		return ""
	}
}
