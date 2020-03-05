package expr

import "github.com/google/nftables/expr"

// Counter creates a counter
func Counter(bytes, packets uint64) []expr.Any {
	// [ counter pkts X bytes Y ]
	return []expr.Any{
		&expr.Counter{
			Bytes:   bytes,
			Packets: packets,
		},
	}
}

// NamedCounter creates a counter
func NamedCounter(name string) []expr.Any {
	// [ objref type 1 name XXX ]
	return []expr.Any{
		&expr.Objref{
			Type: 1,
			Name: name,
		},
	}
}
