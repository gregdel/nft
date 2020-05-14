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

// VerdictAccept accepts the packet
func VerdictAccept() []expr.Any {
	return []expr.Any{
		&expr.Verdict{
			Kind: expr.VerdictAccept,
		},
	}
}

// VerdictDrop drops the packet
func VerdictDrop() []expr.Any {
	return []expr.Any{
		&expr.Verdict{
			Kind: expr.VerdictDrop,
		},
	}
}
