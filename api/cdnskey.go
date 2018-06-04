package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetCdnskeyRecord(record *v1.ResourceRecord) dns.RR {
	Cdnskey := record.GetCdnskey()
	return &dns.CDNSKEY{
		DNSKEY: dns.DNSKEY{
			Hdr: dns.RR_Header{
				Name:   record.Header.Name + ".",
				Rrtype: dns.TypeCDNSKEY,
				Class:  dns.ClassINET,
				Ttl:    0},
			Algorithm: uint8(Cdnskey.Algorithm),
			Flags:     uint16(Cdnskey.Flags),
			Protocol:  uint8(Cdnskey.Protocol),
			PublicKey: Cdnskey.PublicKey,
		},
	}
}
