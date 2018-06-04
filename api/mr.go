package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMrRecord(record *v1.ResourceRecord) dns.RR {
	mr := record.GetMr()
	return &dns.MR{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMR,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mr: mr.Mr,
	}
}
