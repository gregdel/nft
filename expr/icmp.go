package expr

import "github.com/google/nftables/expr"

// ICMPTypeValue represents the ICMP types
type ICMPTypeValue int

// ICMPType values
const (
	ICMPTypeEchoReply              ICMPTypeValue = 0
	ICMPTypeDestinationUnreachable ICMPTypeValue = 3
	ICMPTypeSourceQuench           ICMPTypeValue = 4
	ICMPTypeRedirect               ICMPTypeValue = 5
	ICMPTypeEchoRequest            ICMPTypeValue = 8
	ICMPTypeRouterAdvertisement    ICMPTypeValue = 9
	ICMPTypeRouterSolicitation     ICMPTypeValue = 10
	ICMPTypeTimeExceeded           ICMPTypeValue = 11
	ICMPTypeParameterProblem       ICMPTypeValue = 12
	ICMPTypeTimestampRequest       ICMPTypeValue = 13
	ICMPTypeTimestampReply         ICMPTypeValue = 14
	ICMPTypeInfoRequest            ICMPTypeValue = 15
	ICMPTypeInfoReply              ICMPTypeValue = 16
	ICMPTypeAddressMaskRequest     ICMPTypeValue = 17
	ICMPTypeAddressMaskReply       ICMPTypeValue = 18
)

// ICMPType matches an ICMP type
func ICMPType(t ICMPTypeValue) []expr.Any {
	return []expr.Any{
		// [ payload load 1b @ transport header + 0 => reg 1 ]
		&expr.Payload{
			Base:         expr.PayloadBaseTransportHeader,
			Len:          1,
			Offset:       0,
			DestRegister: 1,
		},
		// [ cmp eq reg 1 0x00000XXX ]
		&expr.Cmp{Op: expr.CmpOpEq, Register: 1, Data: []byte{byte(t)}},
	}
}
