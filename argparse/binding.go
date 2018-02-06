package argparse

type Binding int

const (
	UNKNOWN Binding = iota
	NONE
	LEFT
	RIGHT
)

func (binding Binding) String() string {
	switch binding {
	case UNKNOWN:
		return "UNKNOWN"
	case NONE:
		return "NONE"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return ""
	}
}
