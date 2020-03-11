package expr

import "github.com/google/nftables/expr"

// VerdictReturn returns from the current chain
func VerdictReturn() []expr.Any {
	return []expr.Any{
		&expr.Verdict{
			Kind: expr.VerdictReturn,
		},
	}
}
