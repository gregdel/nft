package expr

import "github.com/google/nftables/expr"

// VerdictReturn returns from the current chain
func VerdictReturn() []expr.Any {
	return []expr.Any{
		&expr.Verdict{
			// This type is not yet implemented, waiting for
			// https://github.com/google/nftables/pull/99
			// Kind: expr.VerdictReturn,
			Kind: expr.VerdictKind(-5),
		},
	}
}
