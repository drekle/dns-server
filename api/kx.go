package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetKxRecord(record *v1.ResourceRecord) dns.RR {
	kx := record.GetKx()
	return &dns.KX{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeKX,
			Class:  dns.ClassINET,
			Ttl:    0},
		Exchanger:  kx.Exchanger,
		Preference: uint16(kx.Preference),
	}
}
