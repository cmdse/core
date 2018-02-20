package parser

import (
	"github.com/cmdse/core/argparse/tkn"
)

// A Behavior is a set of hooks allowing to configure parser's strategy.
// Those hooks are run in Parser#ParseTokens method.
type Behavior struct {
	// Instructions which will be run at the beginning of Parser#ParseTokens
	RunStaticChecks func(p *Parser)
	// Instructions which will be run for each Parser#ParseTokens' pass
	RunInferences func(*Parser, *tkn.Token)
}
