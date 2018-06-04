package api

import (
	"github.com/drekle/dns-server/pkg/lib/go/v1"
	"github.com/miekg/dns"
)

func GetSpfRecord(record *v1.ResourceRecord) dns.RR {
	spf := record.GetSpf()
	return &dns.SPF{
		Hdr: dns.RR_Header{
			Name:   record.Header.Name + ".",
			Rrtype: dns.TypeSPF,
			Class:  dns.ClassINET,
			Ttl:    0},
		Txt: spf.Txt,
	}
}
