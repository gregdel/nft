package expr

import "github.com/google/nftables/expr"

// LookupNamedSet lookup the value in a named set
func LookupNamedSet(name string) []expr.Any {
	// [ lookup reg 1 set gscan ]
	return []expr.Any{
		&expr.Lookup{
			SourceRegister: 1,
			SetName:        name,
		},
	}
}
