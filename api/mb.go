package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetMbRecord(record *v1.ResourceRecord) dns.RR {
	mb := record.GetMb()
	return &dns.MB{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeMB,
			Class:  dns.ClassINET,
			Ttl:    0},
		Mb: mb.Mb,
	}
}
