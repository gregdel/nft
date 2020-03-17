package expr

import (
	"github.com/google/nftables/expr"
)

// Limit reprents a rate limit
func Limit(t expr.LimitType, rate uint64, unit expr.LimitTime, over bool, burst uint32) []expr.Any {
	return []expr.Any{
		&expr.Limit{
			Type:  t,
			Rate:  rate,
			Over:  over,
			Unit:  unit,
			Burst: burst,
		},
	}
}
