package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetL64Record(record *v1.ResourceRecord) dns.RR {
	l64 := record.GetL64()
	return &dns.L64{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeL64,
			Class:  dns.ClassINET,
			Ttl:    0},
		Locator64:  l64.Locator64,
		Preference: uint16(l64.Preference),
	}
}
