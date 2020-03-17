package expr

import (
	"net"

	"github.com/google/nftables/expr"
)

// ip matches the source or destination ip / range
func ip(ipnet *net.IPNet, offset uint32) []expr.Any {
	// TODO: make it work with ipv6
	// TODO: remote the bitwise operation if there's not mask
	if ipnet == nil {
		return nil
	}

	if ipnet.IP.To4() == nil {
		return nil
	}

	// [ payload load 4b @ network header + 12 => reg 1 ]
	// [ bitwise reg 1 = (reg=1 & YYY ) ^ 0x00000000 ]
	// [ cmp eq reg 1 XXX ]
	return []expr.Any{
		&expr.Payload{
			OperationType: expr.PayloadLoad,
			DestRegister:  1,
			Base:          expr.PayloadBaseNetworkHeader,
			Len:           4,
			Offset:        offset,
		},
		&expr.Bitwise{
			SourceRegister: 1,
			DestRegister:   1,
			Len:            4,
			Mask:           ipnet.Mask,
			Xor:            []byte{0x0, 0x0, 0x0, 0x0},
		},
		&expr.Cmp{
			Op:       expr.CmpOpEq,
			Register: 1,
			Data:     ipnet.IP.To4(),
		},
	}
}

// SrcIP matches the source IP
func SrcIP(ipnet *net.IPNet) []expr.Any {
	return ip(ipnet, 12)
}

// DestIP matches the destination IP
func DestIP(ipnet *net.IPNet) []expr.Any {
	return ip(ipnet, 16)
}
