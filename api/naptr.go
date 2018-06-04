package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNaptrRecord(record *v1.ResourceRecord) dns.RR {
	naptr := record.GetNaptr()
	return &dns.NAPTR{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNAPTR,
			Class:  dns.ClassINET,
			Ttl:    0},
		Flags:       naptr.Flags,
		Order:       uint16(naptr.Order),
		Preference:  uint16(naptr.Preference),
		Regexp:      naptr.Regexp,
		Replacement: naptr.Replacement,
		Service:     naptr.Service,
	}
}
