package expr

import (
	"net"

	"github.com/google/nftables/binaryutil"
	"github.com/google/nftables/expr"
	"golang.org/x/sys/unix"
)

// DNATIP creates a DNAT to a given ip
func DNATIP(ip net.IP) []expr.Any {
	// [ immediate reg 1 XXX ]
	// [ nat dnat ip addr_min reg 1 addr_max reg 0 ]
	return []expr.Any{
		&expr.Immediate{
			Register: 1,
			Data:     ip.To4(),
		},
		&expr.NAT{
			Type:       expr.NATTypeDestNAT,
			Family:     unix.NFPROTO_IPV4, // TODO: get the family from the given IP
			RegAddrMin: 1,
		},
	}
}

// DNATIPPort creates a DNAT to a given ip / port
func DNATIPPort(ip net.IP, port uint16) []expr.Any {
	// [ immediate reg 1 XXX ]
	// [ immediate reg 2 XXX ]
	// [ nat dnat ip addr_min reg 1 addr_max reg 0 proto_min reg 2 proto_max reg 0 ]
	return []expr.Any{
		&expr.Immediate{
			Register: 1,
			Data:     ip.To4(),
		},
		&expr.Immediate{
			Register: 2,
			Data:     binaryutil.BigEndian.PutUint16(port),
		},
		&expr.NAT{
			Type:        expr.NATTypeDestNAT,
			Family:      unix.NFPROTO_IPV4, // TODO: get the family from the given IP
			RegAddrMin:  1,
			RegProtoMin: 2,
		},
	}
}
