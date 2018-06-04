package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetRkeyRecord(record *v1.ResourceRecord) dns.RR {
	rkey := record.GetRkey()
	return &dns.RKEY{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeRKEY,
			Class:  dns.ClassINET,
			Ttl:    0},
		Algorithm: uint8(rkey.Algorithm),
		Flags:     uint16(rkey.Flags),
		Protocol:  uint8(rkey.Protocol),
		PublicKey: rkey.PublicKey,
	}
}
