package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetNimlocRecord(record *v1.ResourceRecord) dns.RR {
	nimloc := record.GetNimloc()
	return &dns.NIMLOC{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeNIMLOC,
			Class:  dns.ClassINET,
			Ttl:    0},
		Locator: nimloc.Locator,
	}
}
