package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetKeyRecord(record *v1.ResourceRecord) dns.RR {
	key := record.GetKey()
	return &dns.KEY{
		DNSKEY: dns.DNSKEY{
			Hdr: dns.RR_Header{
				Name:   record.Header.Name + ".",
				Rrtype: dns.TypeKEY,
				Class:  dns.ClassINET,
				Ttl:    0},
			Algorithm: uint8(key.Algorithm),
			Flags:     uint16(key.Flags),
			Protocol:  uint8(key.Protocol),
			PublicKey: key.PublicKey,
		},
	}
}
