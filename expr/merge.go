package expr

import "github.com/google/nftables/expr"

// Merge merges expressions
func Merge(exprs ...[]expr.Any) []expr.Any {
	ret := []expr.Any{}
	for _, e := range exprs {
		ret = append(ret, e...)
	}
	return ret
}
