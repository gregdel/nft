package expr

import (
	"github.com/google/nftables/expr"
)

// L4ProtoValue represents the l4 protocol type
type L4ProtoValue byte

// L4 protocol values
const (
	L4ProtoICMP L4ProtoValue = 0x1
	L4ProtoUDP  L4ProtoValue = 0x11
)

// L4Proto returns an expressions matching the a protocol number
func L4Proto(proto L4ProtoValue) []expr.Any {
	return []expr.Any{
		// [ meta load l4proto => reg 1 ]
		&expr.Meta{Key: expr.MetaKeyL4PROTO, Register: 1},
		// [ cmp eq reg 1 0x00000XXX ]
		&expr.Cmp{Op: expr.CmpOpEq, Register: 1, Data: []byte{byte(proto)}},
	}
}
