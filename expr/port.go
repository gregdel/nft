package expr

import (
	"github.com/google/nftables/binaryutil"
	"github.com/google/nftables/expr"
)

// DPort matches a destination port
func DPort(dport uint16) []expr.Any {
	return []expr.Any{
		// [ payload load 2b @ transport header + 2 => reg 1 ]
		&expr.Payload{
			Base:         expr.PayloadBaseTransportHeader,
			Len:          2,
			Offset:       2,
			DestRegister: 1,
		},
		// [ cmp eq reg 1 0x00003500 ]
		&expr.Cmp{
			Op:       expr.CmpOpEq,
			Register: 1,
			Data:     binaryutil.BigEndian.PutUint16(dport),
		},
	}
}
