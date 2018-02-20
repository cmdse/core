package parser

import (
	"github.com/cmdse/core/argparse/tkn"
)

// A Behavior is a set of hooks allowing to configure parser's strategy.
type Behavior struct {
	// Instructions which will be run before passes
	RunStaticChecks func(p *Parser)
	// Instructions which will be run for each pass
	RunInferences func(*Parser, *tkn.Token)
}
